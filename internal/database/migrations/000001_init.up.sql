CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    email TEXT NOT NULL,
    name TEXT NOT NULL,
    age SMALLINT,
    biography TEXT,
    address TEXT NOT NULL,
    number TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE products (
    id BIGSERIAL PRIMARY KEY,
    users_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    metadata JSONB NOT NULL DEFAULT '{}'::jsonb,
    description TEXT NOT NULL,
    price NUMERIC(10,2) NOT NULL,
    tags TEXT[] DEFAULT ARRAY[]::TEXT[],
    location TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

-- Buscar produtos por usuário
CREATE INDEX idx_products_usuarios ON products(users_id);

-- Buscar produtos por tags (GIN)
CREATE INDEX idx_products_tags ON products USING GIN(tags);

-- Produtos ainda ativos (parcial)
CREATE INDEX idx_products_created_at ON products(created_at);

-- Usuário + Localização (composto)
CREATE INDEX idx_user_location ON products(users_id, location);
