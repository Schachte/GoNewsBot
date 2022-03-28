# Go News Bot

## Background

This bot is designed to pull content dynamically from a given set of `Sources` and upload them in a _fanout_ style across the configured `Sinks`. For example, in this current implementation, the top-voted hacker news post will upload to twitter any time that post changes. 
