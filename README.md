# Kratos Contrib

m/_editor.lua:0: InsertEnter 自动命令 "*"..script nvim_exec2() called at InsertEnter 自动命令 "*":0../Users/xsh/.local/share/nvim/lazy/LuaSnip/plugin/luasnip.lua: Vim(source):E5113: Error while calling lua chunk: ...cal/share/nvim/lazy/LuaSnip/lua/luasnip/loaders/init.lua:1: module 'luasnip.loaders._caches' not found:
	no field package.preload['luasnip.loaders._caches']
cache_loader: module luasnip.loaders._caches not found
cache_loader_lib: module luasnip.loaders._caches not found
	no file './luasnip/loaders/_caches.lua'
	no file '/usr/local/share/luajit-2.1/luasnip/loaders/_caches.lua'
	no file '/usr/local/share/lua/5.1/luasnip/loaders/_caches.lua'
	no file '/usr/local/share/lua/5.1/luasnip/loaders/_caches/init.lua'
	no file './luasnip/loaders/_caches.so'
	no file '/usr/local/lib/lua/5.1/luasnip/loaders/_caches.so'
	no file '/usr/local/lib/lua/5.1/loadall.so'
	no file './luasnip.so'
	no file '/usr/local/lib/lua/5.1/luasnip.so'
	no file '/usr/local/lib/lua/5.1/loadall.so'
stack traceback:
	[C]: in function 'require'
	...cal/share/nvim/lazy/LuaSnip/lua/luasnip/loaders/init.lua:1: in main chunk
	[C]: in function 'require'
	.../xsh/.local/share/nvim/lazy/LuaSnip/lua/luasnip/init.lua:7: in main chunk
	[C]: in function 'require'
	...sh/.local/share/nvim/lazy/LuaSnip/lua/luasnip/config.lua:243: in function '_setup'
	...rs/xsh/.local/share/nvim/lazy/LuaSnip/plugin/luasnip.lua:79: in main chunk
	[C]: in function 'nvim_exec2'
	vim/_editor.lua: in function 'cmd'
	...local/share/nvim/lazy/lazy.nvim/lua/lazy/core/loader.lua:503: in function <...local/share/nvim/lazy/lazy.nvim/lua/lazy/core/loader.lua:502>
	[C]: in function 'xpcall'
	...
	/Users/xsh/.config/nvim/lua/configs/nvim-cmp.lua:1: in main chunk

