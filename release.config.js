module.exports = {
    branches: [{name: 'main'}],
    plugins: [
        "@semantic-release/commit-analyzer",
        "@semantic-release/release-notes-generator",
        "@semantic-release/github"
    ],
    verifyConditions: [
        "@semantic-release/github"
    ],
    publish: [
        "@semantic-release/github"
    ]
}
