import requests

# URL of the word list (DWYL English words)
WORD_LIST_URL = "https://raw.githubusercontent.com/dwyl/english-words/master/words_alpha.txt"
OUTPUT_FILE = "words.txt"

def main():
    # Fetch the word list
    print("Downloading word list...")
    response = requests.get(WORD_LIST_URL)
    response.raise_for_status()
    all_words = response.text.splitlines()

    # Filter for 4-letter words
    words4 = [word.lower() for word in all_words if len(word) == 4]

    # Save to words.txt
    print(f"Writing {len(words4)} four-letter words to {OUTPUT_FILE}...")
    with open(OUTPUT_FILE, "w") as f:
        f.write("\n".join(words4))

    print("Done!")

if __name__ == "__main__":
    main()