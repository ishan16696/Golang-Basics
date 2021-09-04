package stringUtils

func Fullname(firstname , lastname string) (string,int){
	full := firstname + " " + lastname
	return full,len(full)
}
