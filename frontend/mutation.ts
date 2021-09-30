import gql from "graphql-tag";

export const CREATE_PLUGIN = gql`
mutation CreatePlugin($name: String!, $type: NevermorePluginType!) {
  createPlugin(type: $type, name: $name) {
    successful
    plugin {
      id
    }
  }
}
`

export const DELETE_UPLOAD_TOKEN = gql`
mutation DeleteUploadToken($id: ID!) {
  deleteUploadToken(id: $id) {
    successful
  }
}
`

export const CREATE_UPLOAD_TOKEN = gql`
mutation CreateUploadToken($pluginID: ID!) {
  createUploadToken(pluginID: $pluginID)
}
`