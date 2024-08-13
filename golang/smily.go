package main

func CountSmilyFace(text []string) int {
	// TODO : start your code here
	return countSmileys(text)
}

func countSmileys(text []string) int {

	// ถ้าไม่มีข้อมูลให้ return 0
	if len(text) == 0 {
		return 0
	}

	countSmile := 0
	for _, s := range text {
		if len(s) == 2 {
			switch s {
			case ":)", ";)", ":D", ";D":
				countSmile++
			}
		} else if len(s) == 3 {
			switch s {
			case ":-)", ";-)", ":~)", ";~)", ":-D", ";-D", ":~D", ";~D":
				countSmile++
			}
		}
	}

	return countSmile
}
