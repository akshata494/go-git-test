// Akshata's test code to start using the go-git package.
// Also, test code for git pulling non fast forward updates.

package main

/* func main() {
	fmt.Println("Cloning test repo using go-git.")

	url := "https://github.com/akshata494/go-git-test.git"
	directory := "git-merge-test"

	fmt.Printf("git clone %s %s --recursive\n", url, directory)

	r, err := git.PlainClone(directory, false, &git.CloneOptions{
		URL:               url,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})

	if err != nil {
		fmt.Printf("PlainClone error : %s.\n", err)
		os.Exit(1)
	}

	ref, err := r.Head()
	if err != nil {
		fmt.Printf("Retrieving HEAD error : %s.\n", err)
		os.Exit(1)
	}

	commit, err := r.CommitObject(ref.Hash())
	if err != nil {
		fmt.Printf("Retrieving commit object error : %s.\n", err)
		os.Exit(1)
	}

	fmt.Println(commit)
}*/
