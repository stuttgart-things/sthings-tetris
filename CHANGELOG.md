## [1.1.2](https://github.com/stuttgart-things/sthings-tetris/compare/v1.1.1...v1.1.2) (2025-02-25)


### Bug Fixes

* fix/fix-total-count ([a9439c6](https://github.com/stuttgart-things/sthings-tetris/commit/a9439c6bf0437c6bde7f475ae8186350aa683a0d))

## [1.1.1](https://github.com/stuttgart-things/sthings-tetris/compare/v1.1.0...v1.1.1) (2025-02-23)


### Bug Fixes

* fix/release ([2ee7bfc](https://github.com/stuttgart-things/sthings-tetris/commit/2ee7bfcb1ec3edb5daa40e8872fbfee7d669f884))

# [1.1.0](https://github.com/stuttgart-things/sthings-tetris/compare/v1.0.3...v1.1.0) (2025-02-23)


### Features

* feature/enhance-logging ([0ff741c](https://github.com/stuttgart-things/sthings-tetris/commit/0ff741c7d59d0661976067ce7db7a3ab1f5b4cdf))

## [1.0.3](https://github.com/stuttgart-things/sthings-tetris/compare/v1.0.2...v1.0.3) (2025-02-23)


### Bug Fixes

* fix/rework-arch ([30bd248](https://github.com/stuttgart-things/sthings-tetris/commit/30bd24837922b25f794e205e8677a3f876573c34))

## [1.0.2](https://github.com/stuttgart-things/sthings-tetris/compare/v1.0.1...v1.0.2) (2025-02-23)


### Bug Fixes

* fix/build-config ([1f31dd0](https://github.com/stuttgart-things/sthings-tetris/commit/1f31dd093840ab454940117ed74150681b00e80a))

## [1.0.1](https://github.com/stuttgart-things/sthings-tetris/compare/v1.0.0...v1.0.1) (2025-02-23)


### Bug Fixes

* fix/gorelease-config ([c35a55c](https://github.com/stuttgart-things/sthings-tetris/commit/c35a55c30725cd628bedd9e2986b19cc19cae75a))

# 1.0.0 (2025-02-23)


### Bug Fixes

* add empty cell check when trying to rotate ([db01fd2](https://github.com/stuttgart-things/sthings-tetris/commit/db01fd2ea8153f6393623a732d808008cdb773fb))
* allow incrementing several levels at once ([ef1bd42](https://github.com/stuttgart-things/sthings-tetris/commit/ef1bd4207c7c5ceb30ed755c6f59688d25268f63))
* amend hold key to use literal space ([3754c8a](https://github.com/stuttgart-things/sthings-tetris/commit/3754c8a4d307c7c7aec5c4b5cab4ca9d5ea8c702))
* amend offset pause message ([#9](https://github.com/stuttgart-things/sthings-tetris/issues/9)) ([d99f482](https://github.com/stuttgart-things/sthings-tetris/commit/d99f482cd6d1095b4bccda6817749edca806952b))
* amend out of bounds checks ([2e59a83](https://github.com/stuttgart-things/sthings-tetris/commit/2e59a837da2eb3ebfce737fba45eef1fe9baace4))
* amend playfield width ([ca50043](https://github.com/stuttgart-things/sthings-tetris/commit/ca500436cb997a6459b50b5f0e59808d250d697b))
* amend tetrimino rotation coordinates ([e86c2f0](https://github.com/stuttgart-things/sthings-tetris/commit/e86c2f034dd69b70303b3a2da689bddf11f36320))
* amend Tetrimino.CanMoveDown, *.canMoveLeft, *.canMoveRight ([c284c69](https://github.com/stuttgart-things/sthings-tetris/commit/c284c690cd09bf39a8ef7656d974f13b851ea48c)), closes [#1](https://github.com/stuttgart-things/sthings-tetris/issues/1)
* amend Tetrimino.MoveRight to use canMoveRight ([6dcedb2](https://github.com/stuttgart-things/sthings-tetris/commit/6dcedb26577c83525ac0fc49830f228e7dc53b4d))
* call update of game during game mode ([6bffcce](https://github.com/stuttgart-things/sthings-tetris/commit/6bffcce2c3ef11f1f53070c5e2433fe9c7f940b8))
* change clashing quit keybind ([3cc354f](https://github.com/stuttgart-things/sthings-tetris/commit/3cc354f0d961cbe4e30a1e7a0b2c9192a3b9bfb2))
* display bottom 20 cells of playfield ([d957346](https://github.com/stuttgart-things/sthings-tetris/commit/d957346466e4a46ec5dc656376f14681f4432b53))
* explicitly enable cgo for releases ([84d23d6](https://github.com/stuttgart-things/sthings-tetris/commit/84d23d63abc1b54834f544ac944e8701ff8735d4))
* fall calculations ([16b8808](https://github.com/stuttgart-things/sthings-tetris/commit/16b88080880c58b7afd7dc29d5cd440143cfed91))
* game over after hard drop ([24613e3](https://github.com/stuttgart-things/sthings-tetris/commit/24613e3efdbeedd22dfcaf123b0f869fcdf4ef8c))
* get past menu using starter ([141889f](https://github.com/stuttgart-things/sthings-tetris/commit/141889fd872964157ea24cd1e4c745147d36511c))
* handle nil pointers in Tetrimino.Copy method ([1bb8826](https://github.com/stuttgart-things/sthings-tetris/commit/1bb8826102c37b6679c826f13fa565fb77bd072e))
* hold styles should ensure the hold box height is constant ([e01d4f4](https://github.com/stuttgart-things/sthings-tetris/commit/e01d4f4b88cb0f0125bf1533504d97a3b7096626))
* hold UI changing size ([a8316a4](https://github.com/stuttgart-things/sthings-tetris/commit/a8316a42468cf4576423cc86c925765096d4cd6c))
* increase fall speed with level ([66e12a2](https://github.com/stuttgart-things/sthings-tetris/commit/66e12a29846f6f90e7545de84965bcb4858235e9))
* init child on mode switch ([a3af993](https://github.com/stuttgart-things/sthings-tetris/commit/a3af9934cc9dbd317f5cfd4158d96c78ede3adee))
* invert isCellEmpty ([b230c27](https://github.com/stuttgart-things/sthings-tetris/commit/b230c27d51f1c9cf4c78e7d65b52a17b8800c5fc))
* leaderboard newline issue ([e23fdb0](https://github.com/stuttgart-things/sthings-tetris/commit/e23fdb07e5792a8db804766c80d15dd021bfb64a))
* marathon reference receivers ([e5093ce](https://github.com/stuttgart-things/sthings-tetris/commit/e5093ce901c525dfe8203b40238209643b4e4dc2))
* move fullscreen functionality to starter ([3e1198f](https://github.com/stuttgart-things/sthings-tetris/commit/3e1198f0a32a73d0cfbfe9ed8ae1aee21359947a))
* only fill bag when it has space for a complete set of Tetriminos ([16253bb](https://github.com/stuttgart-things/sthings-tetris/commit/16253bb93f17c8031b8c04200c6adbb9e8a09998))
* refill bag when there is room ([a43ec48](https://github.com/stuttgart-things/sthings-tetris/commit/a43ec480fde33ccfd2e1a4692279f5642647e3e7))
* remove CGO disabled in releases ([96bd924](https://github.com/stuttgart-things/sthings-tetris/commit/96bd924f218e4f0e6cb571184ec1735fbf9b0d0f))
* remove debugging log ([23d48e5](https://github.com/stuttgart-things/sthings-tetris/commit/23d48e5255e6e68bac99dafdb6baf7413cd2d594))
* reset hold tetrimino ([643d0de](https://github.com/stuttgart-things/sthings-tetris/commit/643d0de7222d177652d7ca09d8e3d5e574e8e6ef))
* restrict max level ([82d77cf](https://github.com/stuttgart-things/sthings-tetris/commit/82d77cf29596d4c3d8d0dbd61ca57787447c839a))
* rotate counter clockwise was going clockwise ([edb7974](https://github.com/stuttgart-things/sthings-tetris/commit/edb7974775d134d8efbe870451049b95d5927fa7))
* soft drop bugs ([#38](https://github.com/stuttgart-things/sthings-tetris/issues/38)) ([c2d38a9](https://github.com/stuttgart-things/sthings-tetris/commit/c2d38a96f3a48e2230adc59029edd6b7f5554300))
* Tetrimino.Copy was not correctly copying cells or rotation coords ([8982360](https://github.com/stuttgart-things/sthings-tetris/commit/898236084ce3e7361621a6ba6ba2b6d7b3082b46))
* update fall stopwatch interval on tick; bug when adding ghost at skyline ([4acf7f1](https://github.com/stuttgart-things/sthings-tetris/commit/4acf7f1c9bec4cdac6799a782fefbd6c1ae786f9))
* update position & current rotation index after tet rotation ([bb32cf7](https://github.com/stuttgart-things/sthings-tetris/commit/bb32cf745ce9b9dfbf6f0b1d67b00de52791c95e))


### Features

* add 7-bag generation ([b825f2b](https://github.com/stuttgart-things/sthings-tetris/commit/b825f2bab5ea02ebff7964adabb01ac6fc64593e))
* add action switch for correct scoring ([e233635](https://github.com/stuttgart-things/sthings-tetris/commit/e233635be4d4684024b4684568ce89c7fc7ebe93))
* add bag view ([e9eb429](https://github.com/stuttgart-things/sthings-tetris/commit/e9eb429566849134fbe2bd88c33d6d448d727837))
* add bounds checks to Matrix.RemoveTetrimino & Matrix.AddTetrimino ([51f359a](https://github.com/stuttgart-things/sthings-tetris/commit/51f359a4781ad7a6778e5498ce841446a883a990))
* add clockwise rotation ([c9df9a6](https://github.com/stuttgart-things/sthings-tetris/commit/c9df9a628291335c462905fc78dbc35ff0a9e3f7))
* add column indicator style ([4e88366](https://github.com/stuttgart-things/sthings-tetris/commit/4e88366ee47b6aaec516863e018887493ef8f9e9))
* add config features ([ba55b87](https://github.com/stuttgart-things/sthings-tetris/commit/ba55b87aa5c35c5fa42627d8b8eb15004bbad637))
* add config for keymap ([0205406](https://github.com/stuttgart-things/sthings-tetris/commit/020540695c2638613d24407b16e053f673ca47a7))
* add config to model; use config for styling ([d79f990](https://github.com/stuttgart-things/sthings-tetris/commit/d79f9908d3a196b1ab55d39abe907b83084a4f5c))
* add config.toml parsing ([d47b3df](https://github.com/stuttgart-things/sthings-tetris/commit/d47b3df2d58a4873098d14049f4affb14b84e36e))
* add counter-clockwise rotation ([2a77990](https://github.com/stuttgart-things/sthings-tetris/commit/2a7799068cc03160c78760f5a1bafca6bd384b58))
* add error wrapping to tetrimino methods ([3d9dc76](https://github.com/stuttgart-things/sthings-tetris/commit/3d9dc7618e9115b9ebc411af97f807832393f19d))
* add flags for config & database file path ([92d04cd](https://github.com/stuttgart-things/sthings-tetris/commit/92d04cd883a5d1182c6b1b2dcc250cadf1192651))
* add fullscreen flag & functionality ([f56322c](https://github.com/stuttgart-things/sthings-tetris/commit/f56322c8df284ae6fde454f1a7f28b54f72f40b8))
* add game over checks ([28bd21a](https://github.com/stuttgart-things/sthings-tetris/commit/28bd21a60ad39752cccc24bd8ae6103bec1dc84e))
* add ghost style ([92e2a3c](https://github.com/stuttgart-things/sthings-tetris/commit/92e2a3c4d8393f59a6e6225aca8b83e68aedb9f4))
* add ghost tet ([1020f15](https://github.com/stuttgart-things/sthings-tetris/commit/1020f156c871727805c47185f6cd4e90868a545f))
* add hard fall ([ed4d066](https://github.com/stuttgart-things/sthings-tetris/commit/ed4d066e47649f9ffab00fdff2f9a091dd670c2d))
* add hold & information views ([2add0e8](https://github.com/stuttgart-things/sthings-tetris/commit/2add0e803cbd785e9052097f8185e5d77ad56924))
* add holding; move lowering tet to dedicated method ([8f350e3](https://github.com/stuttgart-things/sthings-tetris/commit/8f350e397dbf8445e8d9e217add9d5bbb500b513))
* add level; calculate fall speeds using level ([e0257d4](https://github.com/stuttgart-things/sthings-tetris/commit/e0257d4aba7cef94f81ea11991384d4639e88501))
* add menu ([9d13513](https://github.com/stuttgart-things/sthings-tetris/commit/9d13513a20cfbc96897a2d5446e02fa77a90d243))
* add menu name input for leaderboard ([daac958](https://github.com/stuttgart-things/sthings-tetris/commit/daac958a4a9938512d4d0ac0996f1b6022d858ac))
* add moving sideways ([8030378](https://github.com/stuttgart-things/sthings-tetris/commit/8030378e0be69d2b10016b782385716b87e91f24))
* add option component type ([e07fca8](https://github.com/stuttgart-things/sthings-tetris/commit/e07fca8b169243a2524d8709fc81499cf1aa814d))
* add pause; add points for soft & hard drops ([9919359](https://github.com/stuttgart-things/sthings-tetris/commit/99193595f95d20b916e9ebc9c6fe65cbc4160bfc))
* add play timer ([81dd7cc](https://github.com/stuttgart-things/sthings-tetris/commit/81dd7ccf4f1d5a9951581a1744ad0aabc9d69088))
* add Playfield.NewTetrimino method ([23ed6c6](https://github.com/stuttgart-things/sthings-tetris/commit/23ed6c6bce159d5a029b3bfdc4753af7d18fb076))
* add remainder of key map ([6201eb0](https://github.com/stuttgart-things/sthings-tetris/commit/6201eb0b8466c960181255cd6544af8c271d0817))
* add row indicator to view ([3b940fe](https://github.com/stuttgart-things/sthings-tetris/commit/3b940fe74321ce740f64446c94c8826459a44027))
* add scoring type & scoring switch ([452b301](https://github.com/stuttgart-things/sthings-tetris/commit/452b301daf8698ee650d5ead1fc4877fd2ad90c5))
* add Scoring.AddSoftDrop & Socring.AddHardDrop methods ([549d23b](https://github.com/stuttgart-things/sthings-tetris/commit/549d23b37a19205673d947dc2c75dd142b527941))
* add soft drop toggling ([275ec5a](https://github.com/stuttgart-things/sthings-tetris/commit/275ec5a784762d41de9e6078e7ed318fc66ece9a))
* add startingPositions global variable; remove unused parameter from method ([9cd7518](https://github.com/stuttgart-things/sthings-tetris/commit/9cd7518835a7becdbbd8f9bcc24984edd35960ec))
* add super rotation system (SRS) ([5a35dd4](https://github.com/stuttgart-things/sthings-tetris/commit/5a35dd471d989d2c5b865ea0994983e7db85045d))
* add tetrimino styles; render tetriminos ([8781374](https://github.com/stuttgart-things/sthings-tetris/commit/87813740cd4cd77f55fb63a4580e4f2f33b9acc8))
* add tetrimino type & methods ([541c367](https://github.com/stuttgart-things/sthings-tetris/commit/541c36797de6e1dcdb17ed6641de249caff6b116))
* add Tetrimino.Copy method ([5278fe8](https://github.com/stuttgart-things/sthings-tetris/commit/5278fe8d7fc11cd09fcdca49c5832a8d0b1df041))
* auto select new leaderboard entry ([92c28c8](https://github.com/stuttgart-things/sthings-tetris/commit/92c28c8a28dee33d4a70efe8177ab29b025773f0))
* connect config game ends on max level ([7ad5629](https://github.com/stuttgart-things/sthings-tetris/commit/7ad56293381ba3d294cc81a9a73adfd5ddb283ae))
* connect config max level ([ea02e7c](https://github.com/stuttgart-things/sthings-tetris/commit/ea02e7cfc1deeb97f7e1f369efe9f8dd25a3049c))
* define program style; define key map ([0aa4006](https://github.com/stuttgart-things/sthings-tetris/commit/0aa40066bf08c85142292abe2b9fbe7130782cb8))
* feature/add-release-cofig ([464f02d](https://github.com/stuttgart-things/sthings-tetris/commit/464f02d21aaede594a7c0d7abf0373e5e2310f95))
* feature/add-release-cofig ([176f7be](https://github.com/stuttgart-things/sthings-tetris/commit/176f7be8971c8e77629314fb3e6feb24882f67ee))
* feature/add-release-cofig ([a9b121c](https://github.com/stuttgart-things/sthings-tetris/commit/a9b121cbb75f109b96129e5d2532c364f8d7b794))
* fix/fix-readme ([d7b1cfe](https://github.com/stuttgart-things/sthings-tetris/commit/d7b1cfeee1ae096d7382f27a9dca2f2f6f65f420))
* get config in main; store config in starter; connect ghostEnabled config ([acd470d](https://github.com/stuttgart-things/sthings-tetris/commit/acd470dc3ad9321ebf801198f48c43ea9b73b76e))
* get hpicker working ([4ddfc73](https://github.com/stuttgart-things/sthings-tetris/commit/4ddfc7331275bc2fa526f93eaacb69936d081514))
* handle zero divisor in positiveMod ([dba2297](https://github.com/stuttgart-things/sthings-tetris/commit/dba229797c105b75524601f478fd6eb44293f474))
* implement sprint game mode ([#11](https://github.com/stuttgart-things/sthings-tetris/issues/11)) ([f59b998](https://github.com/stuttgart-things/sthings-tetris/commit/f59b99826f27e50326b6c586030e0fac62794dfa))
* implement ultra game mode ([#10](https://github.com/stuttgart-things/sthings-tetris/issues/10)) ([0b78a5a](https://github.com/stuttgart-things/sthings-tetris/commit/0b78a5a4d6d3c65afd5923a1873d9fa5ba002534))
* initialise module; add program model ([e76c022](https://github.com/stuttgart-things/sthings-tetris/commit/e76c02217753ae886bb7bad2fce8af16f319317a))
* main ([a4b08c3](https://github.com/stuttgart-things/sthings-tetris/commit/a4b08c38ac95e6687bbed36cf9658ef3d8c2cac9))
* main ([c00b906](https://github.com/stuttgart-things/sthings-tetris/commit/c00b90653517823db701f2e7526f808902145b08))
* move actions to type safe enums ([b05e160](https://github.com/stuttgart-things/sthings-tetris/commit/b05e160560aab87334f29ada44fbbafb82e279ee))
* move to using rotation compass ([9dd1bca](https://github.com/stuttgart-things/sthings-tetris/commit/9dd1bcab49735163c27983ed3c00361d12fa21c7))
* only add tet to matrix when no longer in play ([ad5fab5](https://github.com/stuttgart-things/sthings-tetris/commit/ad5fab5dbf2316c6dd9f1917991985599b1ec267))
* overlay paused & game over messages ([36c5ee5](https://github.com/stuttgart-things/sthings-tetris/commit/36c5ee5925e7ff6e157a007bd4c9a2f85fa9a028))
* process scores for leaderboard ([3a89c7c](https://github.com/stuttgart-things/sthings-tetris/commit/3a89c7cf9822a393ce3c04b9bb47f17ee1e68a8f))
* quit on game over ([5e13bc7](https://github.com/stuttgart-things/sthings-tetris/commit/5e13bc7a3bb779265f6bc9bf1106d4adc24e0b46))
* refactor project ([#22](https://github.com/stuttgart-things/sthings-tetris/issues/22)) ([c270928](https://github.com/stuttgart-things/sthings-tetris/commit/c270928b6305622cbfc6d8ff307a0791f2affeca))
* remove game over stopwatch; add overlay message for pause & game over ([fdf6f6e](https://github.com/stuttgart-things/sthings-tetris/commit/fdf6f6e0569feacf6fa34e01ad5cd9e3e8f7f9b4))
* remove line on completion ([47ff36d](https://github.com/stuttgart-things/sthings-tetris/commit/47ff36d42d753d522a1da01f70c06a18f74c6e0a))
* render column indicators & tetrimino ghost ([6c46a17](https://github.com/stuttgart-things/sthings-tetris/commit/6c46a1720a28d1228a339fcf5e125efd71435cc3))
* render help component ([83bc52b](https://github.com/stuttgart-things/sthings-tetris/commit/83bc52b2277bf2eaa2ab0650f6be7034f2ccec6c))
* return to menu from the leaderboard; show leaderboard help ([#5](https://github.com/stuttgart-things/sthings-tetris/issues/5)) ([91d2c34](https://github.com/stuttgart-things/sthings-tetris/commit/91d2c34ddfb24e57a64bf7fde01c6429a2b84fc0))
* save new leaderboard entry on game over ([a0fbd00](https://github.com/stuttgart-things/sthings-tetris/commit/a0fbd0008b7006dc66789405d4cb66541c039525))
* set score rank on DB get ([9ba24ae](https://github.com/stuttgart-things/sthings-tetris/commit/9ba24ae81162c8599030cb5b21df1bed7b3ab7a3))
* set score rank on DB get ([0c73ec6](https://github.com/stuttgart-things/sthings-tetris/commit/0c73ec685794f748de1b459d79eb642d0aa4b1e9))
* show first row right away; reset fall timer after hard drop ([8163b03](https://github.com/stuttgart-things/sthings-tetris/commit/8163b03836962e549aa95733504bc1aafc65c75d))
* show leaderboard on game over ([706b986](https://github.com/stuttgart-things/sthings-tetris/commit/706b98692a973e8c2c498c0648422153db50d582))
* start with random tetrimino; add timed falling ([9a48a79](https://github.com/stuttgart-things/sthings-tetris/commit/9a48a7954ff6005fd5fa557cee0d77d04e1eb4b7))
* use a shared set of keys for all models ([4b0f440](https://github.com/stuttgart-things/sthings-tetris/commit/4b0f44019b425d9ce151599125411dd80824476b))
* use alecthomas/kong to add marathon subcommand to CLI ([42bc75c](https://github.com/stuttgart-things/sthings-tetris/commit/42bc75c963e5f6f5a4eb78c4c7cd4e2151302bb3))
* use XDG for config ([#24](https://github.com/stuttgart-things/sthings-tetris/issues/24)) ([fa956f6](https://github.com/stuttgart-things/sthings-tetris/commit/fa956f6c4b2cdd38ef9bdde6a5eb01d9009445ed))
* use XDG specification ([#23](https://github.com/stuttgart-things/sthings-tetris/issues/23)) ([28565ce](https://github.com/stuttgart-things/sthings-tetris/commit/28565cea319c3b53437db6c7a5a7562be425a82e))
