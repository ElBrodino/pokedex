import sys

from stats import get_num_char, get_num_words


def get_book_text(filepath):
    with open(filepath) as f:
        file_contents = f.read()
        return file_contents


def main():
    # path = "books/frankenstein.txt"
    if len(sys.argv) < 2:
        print("Usage: python3 main.py <path_to_book>")
        sys.exit(1)
    path = sys.argv[1]
    temp = get_book_text(path)

    print("============ BOOKBOT ============")
    print(f"Analyzing book found at {path}...")
    print("----------- Word Count ----------")
    print(f"Found {get_num_words(temp)} total words")
    print("--------- Character Count -------")

    letters = get_num_char(temp)
    for x in letters:
        print(f"{x}: {letters[x]}")


if __name__ == "__main__":
    main()
