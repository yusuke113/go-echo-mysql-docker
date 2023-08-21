-- 既存のデータを削除
DELETE FROM tasks;

INSERT INTO
    tasks (body)
VALUES
    ('John Doe'),
    ('Jane Smith'),
    ('Alice Johnson'),
    ('Bob White');