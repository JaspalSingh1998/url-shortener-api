CREATE TABLE org_members (
    user_id BIGINT NOT NULL,
    org_id BIGINT NOT NULL,
    role TEXT NOT NULL,

    PRIMARY KEY (user_id, org_id),

    CONSTRAINT fk_org_members_user
        FOREIGN KEY (user_id) REFERENCES users (id),

    CONSTRAINT fk_org_members_org
        FOREIGN KEY (org_id) REFERENCES organizations (id)
);
