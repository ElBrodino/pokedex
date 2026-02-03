def get_num_words(text):
    return len(text.split())


def get_num_char(text):
    lower = text.lower()
    # print(f"lower type {type(lower)}")

    split = lower.split()
    # print(f"split type : {type(split)}")

    ordered = []
    for word in split:
        for char in word:
            ordered.append(char)
    ordered.sort()
    # print(f"ordered content : {ordered}")
    # print(f"ordered type : {type(ordered)}")

    dict = {}
    for char in ordered:
        dict.update({char: 0})

    for char in ordered:
        dict[char] += 1
    # print(f"dict content : {dict}")

    return dict
