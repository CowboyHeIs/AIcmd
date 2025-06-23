# **__AI CMD__**

## __*Prerequisite*__

__[ollama](https://ollama.com/download)__

## __*Explanation*__

This is for using local AI in your cmd, You can use any AIs in [ollama](https://ollama.com/download) straight in your Command Prompt either to help yourself especially when you're offline. And yes, if your RAM is big, this may be better than ChatGPT, Claude, Gemini, or DeepSeek (if you picked the best free AI of course).

## __*How to use*__
> __ai__ = _ai {Prompt}_

> __read__ = _read {file}_ (even if {file} is in subdirectory)

> __save__ = _save {ConversationName}_ (Save current _config/files.txt_, _config/debug.txt_, and _config/log.txt_ into __*convo/{ConversationName}*__)

> __load__ = _load {ConversationName}_ (Overwrites current _config/files.txt_, _config/debug.txt_, and _config/log.txt_ from __*convo/{ConversationName}*__)

> __clear__ - _clear_ (Clears current _config/files.txt_, _config/debug.txt_, and _config/log.txt_ )

## __Notes__
- **Change "AI" in ai.py to the model you want to use.**
- **Use __NoDelete.txt__ for files you want to keep in root folder next to ai.bat**
- Use __personality.txt__ to the desired personality you want your AI to have.
- Use __userInfo.txt__ to what your AI need to know about you.
- __log.txt__ saves past conversations, so AI still saves history and you're free to delete or modify as expected.
- Use __debug.txt__ if you want to know what the full prompt looks like.

## __*On Progress*__
- Fixing error messages when using __read.bat__

## __Personal Notes:__
I made this just because ChatGPT still lacks at coding and reasoning, I often use deepseek-coder-v2 or devstral. also, this goes straight to CMD, which makes my coding assignment faster since I can focus on 2 tabs only (IDE and CMD). Also, The code itself are not neat (I'm lazy) so you're free to fix it yourself.



