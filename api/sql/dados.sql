USE devbook;

INSERT INTO usuarios (nome, nick, email, senha) VALUES 
('João Silva', 'joaos', 'joao.silva@example.com', '$2a$10$a8yrh8p7vbRip0j6q5.ySeNA1iTDcrZ1grx/4nujkpcqRCakyrieC'),
('Maria Oliveira', 'mariao', 'maria.oliveira@example.com', '$2a$10$a8yrh8p7vbRip0j6q5.ySeNA1iTDcrZ1grx/4nujkpcqRCakyrieC'),
('Carlos Souza', 'carloss', 'carlos.souza@example.com', '$2a$10$a8yrh8p7vbRip0j6q5.ySeNA1iTDcrZ1grx/4nujkpcqRCakyrieC'),
('Ana Pereira', 'anap', 'ana.pereira@example.com', '$2a$10$a8yrh8p7vbRip0j6q5.ySeNA1iTDcrZ1grx/4nujkpcqRCakyrieC');

INSERT INTO seguidores (usuario_id, seguidor_id) VALUES 
(1, 2),
(1, 3),
(2, 1),
(3, 1),
(4, 2);

INSERT INTO publicacoes (titulo, conteudo, autor_id) VALUES 
('Primeira Publicação', 'Este é o conteúdo da primeira publicação.', 1),
('Segunda Publicação', 'Este é o conteúdo da segunda publicação.', 2),
('Terceira Publicação', 'Este é o conteúdo da terceira publicação.', 3),
('Quarta Publicação', 'Este é o conteúdo da quarta publicação.', 4),
('Quinta Publicação', 'Este é o conteúdo da quinta publicação.', 1);