CREATE TABLE "comments" (
  "id" bigint NOT NULL,
  "text" text NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp
);

CREATE TABLE "post_subreddit" (
  "post_id" bigint NOT NULL,
  "subreddit_id" bigint NOT NULL
);

CREATE TABLE "post_user" (
  "post_id" bigint NOT NULL,
  "user_id" bigint NOT NULL
);

CREATE TABLE "posts" (
  "id" bigint NOT NULL,
  "postname" "character varying" NOT NULL,
  "url" "character varying" NOT NULL,
  "description" text NOT NULL,
  "vote_count" bigint NOT NULL DEFAULT 0,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp,
  "comment_post" bigint,
  "vote_post" bigint
);

CREATE TABLE "subreddits" (
  "id" bigint NOT NULL,
  "name" "character varying" NOT NULL,
  "description" text NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp
);

CREATE TABLE "users" (
  "id" bigint NOT NULL,
  "username" "character varying" NOT NULL,
  "email" "character varying" NOT NULL,
  "password" "character varying" NOT NULL,
  "enabled" boolean NOT NULL DEFAULT false,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp,
  "comment_user" bigint,
  "subreddit_user" bigint,
  "verification_token_user" bigint,
  "vote_user" bigint
);

CREATE TABLE "verification_tokens" (
  "id" bigint NOT NULL,
  "token" "character varying" NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp
);

CREATE TABLE "votes" (
  "id" bigint NOT NULL,
  "vote" "character varying" NOT NULL
);

ALTER TABLE "post_subreddit" ADD CONSTRAINT "post_subreddit_post_id" FOREIGN KEY ("post_id") REFERENCES "posts" ("id") ON DELETE CASCADE;

ALTER TABLE "post_subreddit" ADD CONSTRAINT "post_subreddit_subreddit_id" FOREIGN KEY ("subreddit_id") REFERENCES "subreddits" ("id") ON DELETE CASCADE;

ALTER TABLE "post_user" ADD CONSTRAINT "post_user_post_id" FOREIGN KEY ("post_id") REFERENCES "posts" ("id") ON DELETE CASCADE;

ALTER TABLE "post_user" ADD CONSTRAINT "post_user_user_id" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "posts" ADD CONSTRAINT "posts_comments_post" FOREIGN KEY ("comment_post") REFERENCES "comments" ("id") ON DELETE SET NULL;

ALTER TABLE "posts" ADD CONSTRAINT "posts_votes_post" FOREIGN KEY ("vote_post") REFERENCES "votes" ("id") ON DELETE SET NULL;

ALTER TABLE "users" ADD CONSTRAINT "users_comments_user" FOREIGN KEY ("comment_user") REFERENCES "comments" ("id") ON DELETE SET NULL;

ALTER TABLE "users" ADD CONSTRAINT "users_subreddits_user" FOREIGN KEY ("subreddit_user") REFERENCES "subreddits" ("id") ON DELETE SET NULL;

ALTER TABLE "users" ADD CONSTRAINT "users_verification_tokens_user" FOREIGN KEY ("verification_token_user") REFERENCES "verification_tokens" ("id") ON DELETE SET NULL;

ALTER TABLE "users" ADD CONSTRAINT "users_votes_user" FOREIGN KEY ("vote_user") REFERENCES "votes" ("id") ON DELETE SET NULL;
