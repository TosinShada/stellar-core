package xdr

func MemoText(text string) Memo {
	return Memo{Type: MemoTypeMemoText, Text: &text}
}

func MemoID(id uint64) Memo {
	idObj := Uint64(id)
	return Memo{Type: MemoTypeMemoId, Id: &idObj}
}

func MemoHash(hash Hash) Memo {
	return Memo{Type: MemoTypeMemoHash, Hash: &hash}
}

func MemoRetHash(hash Hash) Memo {
	return Memo{Type: MemoTypeMemoReturn, RetHash: &hash}
}

func MemoText1024B(text string) Memo {
	return Memo{Type: MemoTypeMemoText1024B, Text: &text}
}

func MemoText2048B(text string) Memo {
	return Memo{Type: MemoTypeMemoText2048B, Text: &text}
}

func MemoText4096B(text string) Memo {
	return Memo{Type: MemoTypeMemoText4096B, Text: &text}
}
