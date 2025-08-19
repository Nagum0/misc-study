package main

func addEmailsToQueue(emails []string) chan string {
	out := make(chan string, len(emails))
	
	for _, email := range emails {
			out <- email
		}

	return out
}
