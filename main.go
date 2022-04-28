package main

func main() {
	url := "https://gitlab.com/api/graphql"
	query := `
		query last_projects($n: Int = 5) {
			projects(last:$n) {
				nodes {
					name
					description
					forksCount
				}
			}
		}
	`
	Service(url, query)
}
