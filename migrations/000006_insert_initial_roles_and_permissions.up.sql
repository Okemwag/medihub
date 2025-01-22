-- Insert roles
INSERT INTO roles (name) VALUES
        ('receptionist'),
        ('doctor');

-- Insert permissions
INSERT INTO permissions (name, description) VALUES
       ('patient.create', 'Create new patient records'),
       ('patient.read', 'View patient records'),
       ('patient.update', 'Update patient records'),
       ('patient.delete', 'Delete patient records');

-- Assign permissions to roles
INSERT INTO role_permissions (role_id, permission_id)
SELECT r.id, p.id
FROM roles r
         CROSS JOIN permissions p
WHERE
    (r.name = 'receptionist' AND p.name IN ('patient.create', 'patient.read', 'patient.update', 'patient.delete'))
   OR (r.name = 'doctor' AND p.name IN ('patient.read', 'patient.update'));


