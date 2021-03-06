package main

import (
	"bufio"
	"flag"
	"log"
	"math/rand"
	"muddle/util"
	"os"
	"path/filepath"
	"time"

	"muddle/translation"

	"github.com/joho/godotenv"
	"github.com/watson-developer-cloud/go-sdk/languagetranslatorv3"
)

const version = "2019-01-23"

var save bool
var noHash bool
var savePath string
var iterations int
var langCode string
var outLangCode string
var reportTranslation bool

func init() {
	rand.Seed(time.Now().UTC().UnixNano())

	godotenv.Load("ibm-credentials.env")

	log.SetOutput(os.Stderr)
	log.SetPrefix("::> ")

	flag.BoolVar(&save, "save", false, "Save to path or the current directory")
	flag.BoolVar(&noHash, "no-hash", false, "Disables hashing of the language story")
	flag.StringVar(&savePath, "path", ".", "The path to save to; defaults to the working directory")
	flag.IntVar(&iterations, "iterations", 5, "The minimum translations to complete before attempting to return to english")
	flag.StringVar(&langCode, "language", "en", "The language code of the input (for example, 'en' for English, 'es' for Spanish, etc)")
	flag.StringVar(&outLangCode, "out-language", "", "The language code of the output. If not provided, defaults to the input language.")
	flag.BoolVar(&reportTranslation, "report", false, "Dump the current text after each translation")

	flag.Parse()

	if outLangCode == "" {
		outLangCode = langCode
	}
}

func main() {
	apiKey := os.Getenv("LANGUAGE_TRANSLATOR_IAM_APIKEY")
	url := os.Getenv("LANGUAGE_TRANSLATOR_URL")

	log.Println("Reading input")
	var text []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	log.Println("Connecting to translator")
	languageTranslator, err := languagetranslatorv3.NewLanguageTranslatorV3(
		&languagetranslatorv3.LanguageTranslatorV3Options{
			Version:   version,
			IAMApiKey: apiKey,
			URL:       url,
		})
	if err != nil {
		panic(err)
	}

	prevLang, result := translation.Translate(languageTranslator, &text, &langCode)
	if reportTranslation {
		log.Printf("Translation: %s\n", result)
	}
	langsUsed := make([]string, 2, 12)
	langsUsed[0], langsUsed[1] = langCode, prevLang

	for i := 1; ; i++ {
		prevLang, result = translation.Translate(languageTranslator, &result, &prevLang)
		if reportTranslation {
			log.Printf("Translation: %s\n", result)
		}
		langsUsed = append(langsUsed, prevLang)
		if i >= iterations && prevLang == outLangCode {
			break
		}
	}

	if save {
		var hash string
		if !noHash {
			log.Println("Hashing translations used")
			hash, err = util.HashTranslations(&result)
			if err != nil {
				log.Fatalf("Error hashing translations: %v", err)
			}
			log.Printf("Hashed translations: %s", hash)
		} else {
			log.Println("Skipping hash")
			hash = ""
		}
		var path string
		var err error
		log.Println("Finding save location")
		path, err = filepath.Abs(savePath)
		if err != nil {
			util.Dump(result)
			log.Println("Result dumped to standard output")
			log.Fatalln("There was an error locating the save path")
		}
		var fileName string
		fileInfo, err := os.Stat(path)
		var mode os.FileMode
		if err != nil {
			log.Println("The path's info was not found. Checking the error")
			if !os.IsNotExist(err) {
				util.Dump(result)
				log.Fatalln("An unknown error occured while trying to get the info of the directory. Result dumped to standard output.")
			} else {
				log.Println("The path did not exist. Creating the directory.")
				err := os.MkdirAll(path, os.ModePerm)
				if err != nil {
					util.Dump(result)
					log.Fatalln("Error creating directory. Result dumped to standard output.")
				}
				log.Println("Trying to get info of created directory.")
				fi, err := os.Stat(path)
				if err != nil {
					util.Dump(result)
					log.Fatalln("Error getting info of created directory. Result dumped to standard output.")
				}
				mode = fi.Mode()
			}
		} else {
			mode = fileInfo.Mode()
		}

		if mode.IsDir() {
			fileName = filepath.Join(path, hash+"-thosewords-output.txt")
		} else {
			util.Dump(result)
			log.Fatalln("The provided path must be a directory")
		}

		log.Printf("Saving to %s\n", fileName)
		file, err := os.Create(fileName)
		defer func() {
			err := file.Close()
			if err != nil {
				panic(err)
			}
		}()
		if err != nil {
			util.Dump(result)
			log.Fatalf("There was an error (%v) while creating the save file. Contents have been dumped to standard output.\n", err)
		}
		bytesWritten, err := util.DumpTo(result, file)
		if err != nil {
			log.Fatalf("Error writing to file. %v bytes written", bytesWritten)
		}
	} else {
		util.Dump(result)
	}
}
