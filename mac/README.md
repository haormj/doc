## mac

### brew

加速
```
# 替换 Homebrew-core 源
cd "$(brew --repo)"
git remote set-url origin https://mirrors.ustc.edu.cn/brew.git

# 替换 Homebrew-core 库
cd "$(brew --repo)/Library/Taps/homebrew/homebrew-core"
git remote set-url origin https://mirrors.ustc.edu.cn/homebrew-core.git

# 替换 Homebrew-bottles 二进制镜像
echo 'export HOMEBREW_BOTTLE_DOMAIN=https://mirrors.ustc.edu.cn/homebrew-bottles' >> ~/.zshrc
source ~/.zshrc

# 替换 Homebrew-cask 源
cd "$(brew --repo)/Library/Taps/homebrew/homebrew-cask"
git remote set-url origin https://mirrors.ustc.edu.cn/homebrew-cask.git
```

恢复
```
# 恢复 Homebrew-core 源
cd "$(brew --repo)"
git remote set-url origin https://github.com/Homebrew/brew.git

# 恢复 Homebrew-core 库
cd "$(brew --repo)/Library/Taps/homebrew/homebrew-core"
git remote set-url origin https://github.com/Homebrew/homebrew-core.git

# 删除自定义的 HOMEBREW_BOTTLE_DOMAIN
sed -i '' '/HOMEBREW_BOTTLE_DOMAIN/d' ~/.zshrc
source ~/.zshrc

# 替换 Homebrew-cask 源
cd "$(brew --repo)/Library/Taps/homebrew/homebrew-core"
git remote set-url origin https://github.com/Homebrew/homebrew-cask.git
```
