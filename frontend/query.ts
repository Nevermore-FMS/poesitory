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