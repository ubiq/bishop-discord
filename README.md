# Bishop - NFT Discord Bot

This is a Discord bot for ERC721 NFT collections.

It utilizes modern [Discord Slash Commands](https://canary.discord.com/developers/docs/interactions/slash-commands).

All of the token metadata is retrieved directly from the tokenURI in the smart contract.

![Bishop from Aliens](https://web.archive.org/web/20211110142351im_/https://www.denofgeek.com/wp-content/uploads/2016/06/bishop-main.jpg)

Inspired by: https://github.com/lucid-eleven/nft-discord-bot

# Supported functions

The following functions are currently supported:

## User Commands
### ***/TOKEN_COMMAND*** *[tokenId]*
Retrieves the NFT based on the supplied tokenId, returns an embed message as follows.

# Build instructions

Bishop uses [Govips](https://github.com/davidbyttow/govips) which is a wrapper for the [libvips](https://github.com/libvips/libvips) image processing library. Refer to that repo for build instructions.

There isn't a native Go SVG library which properly handles SVG processing (trust me, I've attempted to use some).
