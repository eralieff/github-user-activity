# 🚀 GitHub Activity Fetcher

A simple Go program that fetches and displays recent GitHub activity for a given user.

[Project Requirements (Technical Specification)](https://roadmap.sh/projects/github-user-activity)

---

## ✨ Features
- Retrieves recent GitHub events for a specified user.
- Supports various event types, including push, issues, stars, forks, and repository creation.
- Formats and displays events in a readable format.

---

## 📥 Installation
Clone the repository:
```sh
  git clone https://github.com/eralieff/github-user-activity.git
  cd github-user-activity
```

---

## 🚀 Usage
Run the program with a GitHub username as an argument:
```sh
  go run main.go <github-username>
```
Example:
```sh
  go run main.go octocat
```

### 📌 Sample Output
```
- Pushed 3 commit(s) to my-repo
- Opened an issue in example-repo
- Starred awesome-project
```

---

## 📜 License
This project is licensed under the MIT License.
