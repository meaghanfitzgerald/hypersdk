//! NFT schema
//! See https://nftschool.dev/reference/metadata-schemas/#ethereum-and-evm-compatible-chains for more information
//! on the ERC-721 NFT metadata schema.

use serde_derive::{Deserialize, Serialize};

#[derive(Serialize, Deserialize, Debug, Default)]
pub struct Nft {
    pub symbol: String,
    pub name: String,
    pub uri: String,
}

impl Nft {
    pub fn with_symbol(mut self, symbol: String) -> Self {
        self.symbol = symbol;
        self
    }

    pub fn with_name(mut self, name: String) -> Self {
        self.name = name;
        self
    }

    pub fn with_uri(mut self, uri: String) -> Self {
        self.uri = uri;
        self
    }
}
