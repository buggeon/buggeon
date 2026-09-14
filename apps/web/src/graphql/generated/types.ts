export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: { input: string; output: string; }
  String: { input: string; output: string; }
  Boolean: { input: boolean; output: boolean; }
  Int: { input: number; output: number; }
  Float: { input: number; output: number; }
  Upload: { input: unknown; output: unknown; }
};

export type Board = {
  __typename?: 'Board';
  cards: Array<Card>;
  cardsStatus: Scalars['String']['output'];
  createdAt: Scalars['String']['output'];
  direction: Scalars['String']['output'];
  id: Scalars['ID']['output'];
  name: Scalars['String']['output'];
  projectId: Scalars['String']['output'];
  themeColor: Scalars['String']['output'];
  updatedAt: Scalars['String']['output'];
};

export type Card = {
  __typename?: 'Card';
  assignees: Array<Member>;
  boardId: Scalars['String']['output'];
  content: Scalars['String']['output'];
  createdAt: Scalars['String']['output'];
  dueDate: Scalars['String']['output'];
  id: Scalars['ID']['output'];
  messages: Array<Message>;
  priority: Scalars['String']['output'];
  status: Scalars['String']['output'];
  title: Scalars['String']['output'];
  updatedAt: Scalars['String']['output'];
};

export type CreateBoardInput = {
  cardsStatus: Scalars['String']['input'];
  direction: Scalars['String']['input'];
  name: Scalars['String']['input'];
  themeColor?: InputMaybe<Scalars['String']['input']>;
};

export type CreateCardInput = {
  content: Scalars['String']['input'];
  dueDate: Scalars['String']['input'];
  priority: Scalars['String']['input'];
  title: Scalars['String']['input'];
};

export type CreateProjectInput = {
  description?: InputMaybe<Scalars['String']['input']>;
  leadId: Scalars['String']['input'];
  members?: InputMaybe<Array<InputMaybe<Scalars['String']['input']>>>;
  name: Scalars['String']['input'];
  progress?: InputMaybe<Scalars['Int']['input']>;
};

export type Member = {
  __typename?: 'Member';
  createdAt: Scalars['String']['output'];
  directions: Array<Scalars['String']['output']>;
  id: Scalars['ID']['output'];
  projectId: Scalars['String']['output'];
  role: Scalars['String']['output'];
  user: User;
};

export type Message = {
  __typename?: 'Message';
  cardId: Scalars['String']['output'];
  content: Scalars['String']['output'];
  createdAt: Scalars['String']['output'];
  id: Scalars['ID']['output'];
  replies: Array<Message>;
  replyTo?: Maybe<Message>;
  sender: User;
  updatedAt: Scalars['String']['output'];
};

export type Mutation = {
  __typename?: 'Mutation';
  createBoard: Board;
  createCard: Card;
  createProject: Project;
  deleteBoard: Scalars['Boolean']['output'];
  deleteCard: Scalars['Boolean']['output'];
  deleteProject: Scalars['Boolean']['output'];
  updateBoard: Board;
  updateCard: Card;
  updateProject: Project;
};


export type MutationCreateBoardArgs = {
  input: CreateBoardInput;
  projectId: Scalars['ID']['input'];
};


export type MutationCreateCardArgs = {
  boardId: Scalars['ID']['input'];
  input: CreateCardInput;
};


export type MutationCreateProjectArgs = {
  input: CreateProjectInput;
};


export type MutationDeleteBoardArgs = {
  boardId: Scalars['ID']['input'];
  projectId: Scalars['ID']['input'];
};


export type MutationDeleteCardArgs = {
  boardId: Scalars['ID']['input'];
  cardId: Scalars['ID']['input'];
};


export type MutationDeleteProjectArgs = {
  projectId: Scalars['ID']['input'];
};


export type MutationUpdateBoardArgs = {
  boardId: Scalars['ID']['input'];
  input: UpdateBoardInput;
};


export type MutationUpdateCardArgs = {
  cardId: Scalars['ID']['input'];
  input: UpdateCardInput;
};


export type MutationUpdateProjectArgs = {
  input: UpdateProjectInput;
  projectId: Scalars['ID']['input'];
};

export type Project = {
  __typename?: 'Project';
  boards: Array<Board>;
  createdAt: Scalars['String']['output'];
  description: Scalars['String']['output'];
  id: Scalars['ID']['output'];
  lead: Member;
  logoUrl: Scalars['String']['output'];
  members: Array<Member>;
  name: Scalars['String']['output'];
  progress: Scalars['Int']['output'];
  updatedAt: Scalars['String']['output'];
};

export type Query = {
  __typename?: 'Query';
  board: Board;
  boards: Array<Board>;
  card: Card;
  cards: Array<Card>;
  member: Member;
  members: Array<Member>;
  message: Message;
  messages: Array<Message>;
  project: Project;
  projects: Array<Project>;
  user: User;
};


export type QueryBoardArgs = {
  boardId: Scalars['ID']['input'];
};


export type QueryBoardsArgs = {
  projectId: Scalars['ID']['input'];
};


export type QueryCardArgs = {
  cardId: Scalars['ID']['input'];
};


export type QueryCardsArgs = {
  boardId: Scalars['ID']['input'];
};


export type QueryMemberArgs = {
  memberId: Scalars['ID']['input'];
};


export type QueryMembersArgs = {
  projectId: Scalars['ID']['input'];
};


export type QueryMessageArgs = {
  messageId: Scalars['ID']['input'];
};


export type QueryMessagesArgs = {
  cardId: Scalars['ID']['input'];
};


export type QueryProjectArgs = {
  projectId: Scalars['ID']['input'];
};


export type QueryProjectsArgs = {
  userId: Scalars['ID']['input'];
};


export type QueryUserArgs = {
  userId: Scalars['ID']['input'];
};

export type Schema = {
  __typename?: 'Schema';
  author: Member;
  createdAt: Scalars['String']['output'];
  direction?: Maybe<Scalars['String']['output']>;
  id: Scalars['ID']['output'];
  name: Scalars['String']['output'];
  updatedAt: Scalars['String']['output'];
  url: Scalars['String']['output'];
};

export type UpdateBoardInput = {
  cardsStatus?: InputMaybe<Scalars['String']['input']>;
  direction?: InputMaybe<Scalars['String']['input']>;
  name?: InputMaybe<Scalars['String']['input']>;
  themeColor?: InputMaybe<Scalars['String']['input']>;
};

export type UpdateCardInput = {
  assignees?: InputMaybe<Array<Scalars['ID']['input']>>;
  content?: InputMaybe<Scalars['String']['input']>;
  dueDate?: InputMaybe<Scalars['String']['input']>;
  priority?: InputMaybe<Scalars['String']['input']>;
  status?: InputMaybe<Scalars['String']['input']>;
  title?: InputMaybe<Scalars['String']['input']>;
};

export type UpdateProjectInput = {
  description?: InputMaybe<Scalars['String']['input']>;
  logoUrl?: InputMaybe<Scalars['String']['input']>;
  name?: InputMaybe<Scalars['String']['input']>;
  progress?: InputMaybe<Scalars['Int']['input']>;
};

export type User = {
  __typename?: 'User';
  avatarUrl: Scalars['String']['output'];
  createdAt: Scalars['String']['output'];
  email: Scalars['String']['output'];
  id: Scalars['ID']['output'];
  login: Scalars['String']['output'];
  name: Scalars['String']['output'];
  role: Scalars['String']['output'];
};
