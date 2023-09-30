# Contributing to IPTABLES-VIZ

Thank you for your interest in contributing to the iptables Visualization Tool project. 

We welcome contributions from the community and appreciate your help in making this tool more powerful and user-friendly, your contributions are valuable.

## Getting Started

Before you start contributing, please take a moment to read through this guide to understand the project's goals, coding standards, and contribution process.

## Getting Help

If you have a question about IPTABLES-VIZ, you can [create an issue](https://github.com/iptables-viz/iptables-viz/issues) in this repo.

### Project Overview

This project aims to create a simple yet scalable iptables visualization tool that can be used both in Kubernetes clusters and traditional Linux environments. It helps users visualize and manage iptables rules effectively.

### How to Contribute

1. **Fork the Repository**: Click the "Fork" button in the upper right corner of the GitHub repository to create a copy of the project in your own GitHub account.

2. **Clone Your Fork**: Clone your forked repository to your local development environment and start making your changes.

   ```
   git clone https://github.com/johndoe/iptables-viz.git
   ```

### Pull Request

1. Create a feature branch from the main branch of your forked repository

   ```
   git checkout -b my-new-feature
   ```
   
2. Make your changes

3. Commit your changes with descriptive commit messages:

   ```
   git commit -s -m "Add a new feature: [Feature Name]"
   ```

4. Push your changes to your forked repository:

   ```
   git push origin my-new-feature
   ```

5. Open a [pull request](https://github.com/26tanishabanik/iptables-viz/pulls) to the main branch of the project's repository.

## Developer Certificate of Origin: Signing your work

### Every commit needs to be signed

The Developer Certificate of Origin (DCO) is a lightweight way for contributors to certify that they wrote or otherwise have the right to submit the code they are contributing to the project.

Contributors sign-off that they adhere to these requirements by adding a `Signed-off-by` line to commit messages.

```
This is my commit message

Signed-off-by: John Doe <johndoe@example.com>
```
Git even has a `-s` command line option to append this automatically to your commit message:
```
$ git commit -s -m 'commit message'
```

Each Pull Request is checked  whether or not commits in a Pull Request do contain a valid Signed-off-by line.

### In case, you didn't sign the commit, you can easily follow the below steps:

```
git checkout <branch-name>
git reset $(git merge-base main <branch-name>)
git add -A
git commit -sm "one commit on <branch-name>"
git push --force
```