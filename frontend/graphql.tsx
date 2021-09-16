import { gql } from '@apollo/client';
export type Maybe<T> = T | null;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
};

export type MutatePluginPayload = {
  __typename?: 'MutatePluginPayload';
  plugin?: Maybe<NevermorePlugin>;
  successful: Scalars['Boolean'];
};

export type Mutation = {
  __typename?: 'Mutation';
  createPlugin?: Maybe<MutatePluginPayload>;
  createUploadToken?: Maybe<Scalars['String']>;
  deleteUploadToken?: Maybe<MutatePluginPayload>;
  uploadPluginVersion?: Maybe<UploadPayload>;
};


export type MutationCreatePluginArgs = {
  name: Scalars['String'];
  type: NevermorePluginType;
};


export type MutationCreateUploadTokenArgs = {
  pluginID: Scalars['ID'];
};


export type MutationDeleteUploadTokenArgs = {
  id: Scalars['ID'];
};


export type MutationUploadPluginVersionArgs = {
  channel: Scalars['String'];
  id: Scalars['ID'];
  version: Scalars['String'];
};

export type NevermorePlugin = {
  __typename?: 'NevermorePlugin';
  channels?: Maybe<Array<NevermorePluginChannel>>;
  id: Scalars['ID'];
  latestFullIdentifier?: Maybe<Scalars['String']>;
  latestVersion?: Maybe<NevermorePluginVersion>;
  name: Scalars['String'];
  owner?: Maybe<User>;
  type: NevermorePluginType;
  uploadTokens?: Maybe<Array<UploadToken>>;
};

export type NevermorePluginChannel = {
  __typename?: 'NevermorePluginChannel';
  name: Scalars['String'];
  plugin?: Maybe<NevermorePlugin>;
  /** Note: Will only show the latest 50 versions - earlier versions can still be requested via pluginVersion(identifier) */
  versions?: Maybe<Array<NevermorePluginVersion>>;
};

export type NevermorePluginPage = {
  __typename?: 'NevermorePluginPage';
  hasNext: Scalars['Boolean'];
  pageNum: Scalars['Int'];
  plugins?: Maybe<Array<NevermorePlugin>>;
};

export enum NevermorePluginType {
  Game = 'GAME',
  Generic = 'GENERIC',
  NetworkConfigurator = 'NETWORK_CONFIGURATOR'
}

export type NevermorePluginVersion = {
  __typename?: 'NevermorePluginVersion';
  channel?: Maybe<NevermorePluginChannel>;
  /** Ephemeral - do not store - always request new url before downloading */
  downloadUrl: Scalars['String'];
  fullIdentifier: Scalars['String'];
  id: Scalars['ID'];
  plugin?: Maybe<NevermorePlugin>;
  readme?: Maybe<Scalars['String']>;
  shortIdentifier: Scalars['String'];
};

export type Query = {
  __typename?: 'Query';
  me?: Maybe<User>;
  plugin?: Maybe<NevermorePlugin>;
  pluginVersion?: Maybe<NevermorePluginVersion>;
  searchPlugins?: Maybe<NevermorePluginPage>;
  user?: Maybe<User>;
};


export type QueryPluginArgs = {
  id?: Maybe<Scalars['ID']>;
  name?: Maybe<Scalars['String']>;
};


export type QueryPluginVersionArgs = {
  versionIdentifier: Scalars['String'];
};


export type QuerySearchPluginsArgs = {
  owner?: Maybe<Scalars['ID']>;
  page?: Maybe<Scalars['Int']>;
  search?: Maybe<Scalars['String']>;
  type?: Maybe<NevermorePluginType>;
};


export type QueryUserArgs = {
  id: Scalars['ID'];
};

export type UploadPayload = {
  __typename?: 'UploadPayload';
  url: Scalars['String'];
};

export type UploadToken = {
  __typename?: 'UploadToken';
  createdAt: Scalars['Int'];
  id: Scalars['ID'];
};

export type User = {
  __typename?: 'User';
  id: Scalars['ID'];
  ownedPlugins?: Maybe<NevermorePluginPage>;
  username: Scalars['String'];
};


export type UserOwnedPluginsArgs = {
  page?: Maybe<Scalars['Int']>;
};
