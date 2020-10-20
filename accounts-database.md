* [ ] tables
    * [ ] Users (with a select only user and a insert only user)
        * [ ] id int not null auto_increment primary key (This is the account node id)
        * [ ] name text not null
        * [ ] username text not null
        * [ ] password text not null - sha3_256
        * [ ] account_type int not null
        * [ ] email text not null
        * [ ] phone_number text not null
    * [ ] Teams (with a select only user and insert only user)
        * [ ] id int not null auto_increment primary key (This is the team node id)
        * [ ] owner_id int not null (This is the contractor account owner id)
        * [ ] active bool not null (This is the state of the team)
        * [ ] team_name text not null
