# LeetCode Solutions Portfolio

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

This repository contains my solutions to LeetCode problems, implemented in **PHP** and **Go**. The primary goals of this project are:

1. Self-education in algorithms and data structures
2. Improving problem-solving skills
3. Demonstrating coding abilities in multiple languages
4. Building a technical portfolio


## ⚠️ Important Note
- All problem statements and original descriptions are property of LeetCode
- Solutions are my own implementations
- This repository is for educational/portfolio purposes only
- Contains spoilers - solve problems yourself first!


## Repository Structure
Each problem has its own directory with the following structure:
```
problems/
└──problem-name/            # Problem slug (e.g., two-sum)
    ├── Solution.php        # PHP implementation
    ├── SolutionTest.php    # PHP tests (PHPUnit)
    ├── solution.go         # Go implementation
    ├── solution_test.go    # Go tests (standard/testing or testify)
    └── README.md           # Problem link and additional notes
```


## Key Features
- **Bilingual implementations**: All solutions provided in PHP and Go
- **Test coverage**: Each solution includes test cases
- **Complexity analysis**: Time and space complexity in code comments
- **Cross-language comparison**: See different language approaches to same problems


## Running Tests

### PHP Tests
1. Install dependencies:
```bash
composer install
```

2. Run tests for specific problem:
```bash
composer test problems/two-sum/SolutionTest.php
```

3. Run all tests:
```bash
composer test:all
```

### Go Tests
1. Install testify:
```bash
go get github.com/stretchr/testify
```

2. Run tests for specific problem:
```bash
go test ./problems/two-sum
```

3. Run all tests:
```bash
go test ./...
```


## 📊 Progress Tracking

| ID  | Problem | Difficulty | PHP | Go | Link |
|-----|---------|------------|-----|----|------|
| 1 | Two Sum | Easy | ✅ | ✅ | [View](problems/two-sum) |
| ... | ... | ... | ... | ... | ... |

**Status Legend:**  
✅ - Complete with tests  
🟡 - Partial implementation  
❌ - Not started  


## 📜 License

The source code in this repository is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

**Important Notes:**
- Problem statements and descriptions are copyrighted by [LeetCode](https://leetcode.com) and not covered by this license
- This repository contains personal solutions for educational purposes only
- All LeetCode problem links reference original content on leetcode.com
