
CREATE TABLE aws_services (
  id serial primary key,
  account_id text,
  created_at timestamptz DEFAULT current_timestamp,
  updated_at timestamptz DEFAULT NULL,
  name text DEFAULT NULL,
  description text DEFAULT NULL
);

CREATE TRIGGER aws_services_updated_at
  BEFORE UPDATE OR INSERT ON aws_services
  FOR EACH ROW
  EXECUTE PROCEDURE set_updated_at();

ALTER TABLE aws_rds_instances ADD COLUMN aws_service_id int REFERENCES aws_services(id) ON DELETE CASCADE;
