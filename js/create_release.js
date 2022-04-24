const { graphql } = require("@octokit/graphql");
// or: import { graphql } from "@octokit/graphql";

const { repository } = await graphql(
    `
    {
      repository(owner: "octokit", name: "graphql.js") {
        issues(last: 3) {
          edges {
            node {
              title
            }
          }
        }
      }
    }
  `,
    {
        headers: {
            authorization: `token secret123`,
        },
    }
);