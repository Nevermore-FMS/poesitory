import gql from "graphql-tag";

export const SEARCH_PLUGINS = gql`
query SearchPlugins($search: String, $page: Int) {
  searchPlugins(search: $search, page: $page) {
    hasNext
    plugins {
      id
      name
      owner {
        id
        username
      }
      type
      latestFullIdentifier
    }
  }
}
`

export const GET_PLUGIN_VERSION = gql`
query GetPluginVersion($versionIdentifier: String!) {
  pluginVersion(versionIdentifier: $versionIdentifier) {
    id
    plugin {
      id
      name
      type
      owner {
        id
        username
      }
      channels {
        name
        versions {
          id
          fullIdentifier
        }
      }
    }
    fullIdentifier
    readme
  }
}
`

export const GET_USER_PLUGINS = gql`
query GetUserPlugins($id: ID!, $page: Int) {
  user(id: $id) {
    id
    username
    ownedPlugins(page: $page) {
      hasNext
      plugins {
        id
        name
        latestFullIdentifier
        type
      }
    }
  }
}
`

export const GET_ME_USERNAME = gql`
query GetMeUsername {
  me {
    id
    username
  }
}
`

export const GET_ME_PLUGINS = gql`
query GetMePlugins($page: Int) {
  me {
    username
    ownedPlugins(page: $page) {
      hasNext
      pageNum
      plugins {
        id
        name
        type
        latestFullIdentifier
      }
    }
  }
}
`

export const GET_PLUGIN = gql`
query GetPlugin($id: ID!) {
  plugin(id: $id) {
    id
    name
    type
    latestFullIdentifier
    owner {
      id
    }
    uploadTokens {
      id
      createdAt
    }
  }
}
`