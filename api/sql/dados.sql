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