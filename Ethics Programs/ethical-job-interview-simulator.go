// This program, the "Ethical Job Interview Simulator," simulates a job interview and its outcome based on whether the applicant is considered ethical or not. It underscores the importance of ethical qualities in the hiring process.
package main

import "fmt"

type JobApplicant struct {
	Name        string
	Ethical     bool
	Interviewer string
	Outcome     string
}

func main() {
	applicant := JobApplicant{Name: "John", Ethical: true, Interviewer: "Jane"}
	fmt.Println("Welcome to the Ethical Job Interview Simulator")

	simulateJobInterview(&applicant)

	fmt.Printf("Interview outcome for %s with %s: %s\n", applicant.Name, applicant.Interviewer, applicant.Outcome)
}

func simulateJobInterview(applicant *JobApplicant) {
	fmt.Printf("Interviewer %s is conducting an interview with %s...\n", applicant.Interviewer, applicant.Name)

	if applicant.Ethical {
		applicant.Outcome = "Successful(Ethical Candidate)"
	} else {
		applicant.Outcome = "Unsuccessful (Not an Ethical Candidate)"
	}
}
