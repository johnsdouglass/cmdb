
CREATE TABLE vaults
(
  id text primary key,
  account_id text,
  created_at timestamptz DEFAULT current_timestamp,
  updated_at timestamptz DEFAULT NULL,
  name text DEFAULT NULL,
  description text DEFAULT NULL
);

CREATE TRIGGER vaults_updated_at
  BEFORE UPDATE OR INSERT ON vaults
  FOR EACH ROW
  EXECUTE PROCEDURE set_updated_at();

