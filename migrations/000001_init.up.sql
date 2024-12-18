CREATE TABLE IF NOT EXISTS applicants (
    email varchar(255) PRIMARY KEY,
    applicant_name varchar(255) NOT NULL,
    created_at timestamp WITH time zone NOT NULL DEFAULT NOW(),
    token uuid UNIQUE NOT NULL,
    prompt text NOT NULL,
    solution text NOT NULL
);

CREATE TABLE IF NOT EXISTS submissions (
    submission_id uuid PRIMARY KEY,
    token uuid NOT NULL REFERENCES applicants (token) ON DELETE CASCADE,
    wrong smallint NOT NULL,
    submission_time timestamp WITH time zone NOT NULL DEFAULT NOW()
);
