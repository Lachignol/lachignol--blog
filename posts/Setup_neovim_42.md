+++
title = "Setup Neovim in 42 School"
description = "Tuto d'installation pour un setup rapide et minimum a l'ecole 42"
date = 2024-04-01

[author]
name = "La Chignol"
email = "Pas d'email for you "

[footer]
copyright = "©LaChignol"

+++

# Setup minimum de Neovim a 42.

## 1-Telecharger le fichier image de neovim.


```bash
curl -LO https://github.com/neovim/neovim/releases/latest/download/nvim.appimage
chmod u+x nvim.appimage
./nvim.appimage
```


## 2-Faire que neovim soit accessible depuis partout sur votre ordinateur.
Ajouter a votre fichier .zshrc l'alias qui pointe ver l'executable nvim afin de pouvoir le lancer depuis n'importe.


```zshrc
alias nvim=/mnt/nfs/homes/votre_login_42/nvim
```

## 3-Telecharger la distrib lazyvim.

- Installer la distrib au lien suivre les etapes presente dans ce lien:
 https://www.lazyvim.org/installation
- Afin de voir toutes les icones il est preferable (pas obligatoire )de telecharger une nerdfont (police) disponible au lien suivant :
  https://www.nerdfonts.com/font-downloads

## 4-definir les tabulation comme reel.

- Ajouter dans le fichier ~/.config/nvim/lua/config/keymaps.lua ceci:
(pour info les noms de dossier sont a chaque fois ecrit en haut du code que je met)
```lua
-- ~/.config/nvim/lua/config/keymaps.lua
local function map(mode, lhs, rhs, opts)
  local options = { noremap = true, silent = true }
  if opts then
    options = vim.tbl_extend("force", options, opts)
  end
  vim.api.nvim_set_keymap(mode, lhs, rhs, options)
end

-- Utiliser des tabulations réelles (physiques) pour la touche Tab
vim.opt.expandtab = false -- Ne pas remplacer les tabulations par des espaces

map("i", "kk", "<Esc>")
map("n", "<leader>tk", "<C-w>t<C-w>K") -- change vertical to horizontal
map("n", "<leader>th", "<C-w>t<C-w>H") -- change horizontal to vertical

-- Move around splits using Ctrl + {h,j,k,l}
map("n", "<C-h>", "<C-w>h")
map("n", "<C-j>", "<C-w>j")
map("n", "<C-k>", "<C-w>k")
map("n", "<C-l>", "<C-w>l")
-- Désactiver les touches fléchées dans Neovim si vous voulez vous entrainer dans le dur decommenter :)
--[[
vim.api.nvim_set_keymap('n', '<Up>', '', { noremap = true, silent = true })
vim.api.nvim_set_keymap('n', '<Down>', '', { noremap = true, silent = true })
vim.api.nvim_set_keymap('n', '<Left>', '', { noremap = true, silent = true })
vim.api.nvim_set_keymap('n', '<Right>', '', { noremap = true, silent = true })

vim.api.nvim_set_keymap('i', '<Up>', '', { noremap = true, silent = true })
vim.api.nvim_set_keymap('i', '<Down>', '', { noremap = true, silent = true })
vim.api.nvim_set_keymap('i', '<Left>', '', { noremap = true, silent = true })
vim.api.nvim_set_keymap('i', '<Right>', '', { noremap = true, silent = true })
]]--

```

## 5-Installer les differents plugins.

- Le chemin des fichiers a cree sont en haut de chaque code:
(Ex:Cree le fichier ft_nvim.lua dans ~/.config/nvim/lua/plugins/ et ajouter le code suivant)

### Plugin pour avoir le header de 42 :
```lua
-- ~/.config/nvim/lua/plugins/ft_nvim.lua
return {
  "vinicius507/ft_nvim",
  cmd = { "FtHeader", "Norme" },
  ft = { "c", "cpp" },
  config = function()
    require("ft_nvim").setup({
      header = {
        enable = true,
        username = "ascordil",
        email = "marvin@42.fr",
      },
      norminette = {
        enable = true,
        cmd = "norminette",
        condition = function()
          return true
        end,
      },
    })
  end,
}
```

### Plugin pour avoir un terminal accessible rapidement :
```lua
-- ~/.config/nvim/lua/plugins/toggleterm.lua
return {
  "akinsho/toggleterm.nvim",
  version = "*", -- Assurez-vous d'utiliser la dernière version
  config = function()
    require("toggleterm").setup({
      size = 20, -- Taille du terminal
      open_mapping = [[<c-\>]], -- Raccourci pour ouvrir le terminal
      direction = "horizontal", -- Direction du terminal (horizontal, vertical, float)
      shell = "/bin/zsh", -- Utiliser zsh comme shell
      close_on_exit = true, -- Fermer le terminal à la sortie
      start_in_insert = true, -- Commencer en mode insertion
      insert_mappings = true, -- Activer les mappages en mode insertion
      terminal_mappings = true, -- Activer les mappages pour le terminal
      highlights = {
        Normal = {
          guifg = "#ffffff", -- Couleur de texte
          guibg = "#1e1e1e", -- Couleur de fond
        },
        NormalFloat = {
          link = "Normal", -- Lien vers Normal pour les fenêtres flottantes
        },
      },
    })
  end,
}
```

### Plugin pour avoir la norminette :
```lua
-- ~/.config/nvim/lua/plugins/norminette42.lua
return {
  "hardyrafael17/norminette42.nvim",
  config = function()
    local norminette = require("norminette")
    norminette.setup({
      runOnSave = true,
      maxErrorsToShow = 5,
      active = true,
    })
  end,
}
```

### Plugin pour commenter rapidement avec les touches gcc:
```lua
-- ~/.config/nvim/lua/plugins/comment.lua
return {
  "numToStr/Comment.nvim",
  version = "*", -- Assurez-vous d'utiliser la dernière version
  config = function()
    require("Comment").setup({
      mappings = {
        basic = true, -- Active les mappages de base pour commenter/décommenter
        extra = true, -- Active les mappages supplémentaires
      },
      pre_hook = function(ctx)
        local U = require("Comment.utils")

        -- Si la ligne est vide, ne pas commenter
        if ctx.ctype == U.ctype.line then
          return require("ts_context_commentstring.utils").get_cursor_context()
        end
      end,
    })
  end,
}
```
## Ces plugins sont optionnels mais donnent plus de style:

### colorscheme:
```lua
-- ~/.config/nvim/lua/plugins/colorscheme.lua
return {


  -- add nord
  { "shaunsingh/nord.nvim" },
  -- add gruvbox
  { "ellisonleao/gruvbox.nvim" },

  -- tokyonight
  {
    "folke/tokyonight.nvim",
    lazy = true,
    opts = { style = "moon" },
  },

  -- catppuccin
  {
    "catppuccin/nvim",
    lazy = true,
    name = "catppuccin",
    opts = {
      integrations = {
        aerial = true,
        alpha = true,
        cmp = true,
        dashboard = true,
        flash = true,
        fzf = true,
        grug_far = true,
        gitsigns = true,
        headlines = true,
        illuminate = true,
        indent_blankline = { enabled = true },
        leap = true,
        lsp_trouble = true,
        mason = true,
        markdown = true,
        mini = true,
        native_lsp = {
          enabled = true,
          underlines = {
            errors = { "undercurl" },
            hints = { "undercurl" },
            warnings = { "undercurl" },
            information = { "undercurl" },
          },
        },
        navic = { enabled = true, custom_bg = "lualine" },
        neotest = true,
        neotree = true,
        noice = true,
        notify = true,
        semantic_tokens = true,
        snacks = true,
        telescope = true,
        treesitter = true,
        treesitter_context = true,
        which_key = true,
      },
    },
    specs = {
      {
        "akinsho/bufferline.nvim",
        optional = true,
        opts = function(_, opts)
          if (vim.g.colors_name or ""):find("catppuccin") then
            opts.highlights = require("catppuccin.groups.integrations.bufferline").get()
          end
        end,
      },
    },
  },
}
```

### Plugin pour avoir le fond de neovim transparent (ajouter une image d'arriere plan dans votre terminal ):
```lua
-- ~/.config/nvim/lua/plugins/transparent.lua
return {
        { "xiyaowong/transparent.nvim", config = function()
            require('transparent').setup({
                enable = true, -- Activer la transparence
                extra_groups = { -- Ajouter des groupes supplémentaires si nécessaire
                    'Normal', 'NormalNC', 'Comment', 'SpecialComment',
                    'LineNr', 'SignColumn', 'EndOfBuffer'
                },
            })
        end },
}
```
