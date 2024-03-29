alphabet = set('abcdefghijklmnopqrstuvwxyz')

def is_pangram(sentence: str) -> bool:
    chars_used = set(filter(lambda x: x in alphabet, sentence.lower()))
    return len(chars_used) == len(alphabet)
