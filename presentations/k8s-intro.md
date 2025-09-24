---
marp: true
math: mathjax
theme: rose-pine
# theme: rose-pine-dawn
# theme: rose-pine-moon
paginate: true
# header: This is a header
# footer: This is a footer
title: Marp custom themes
---

<style lang=css>
/*
Rosé Pine theme create by RAINBOWFLESH
> www.rosepinetheme.com

palette in :root
*/

@import "default";
@import "schema";
@import "structure";

:root {
  --base: #232136;
    --surface: #2a273f;
    --overlay: #393552;
    --muted: #6e6a86;
    --subtle: #908caa;
    --text: #e0def4;
    --love: #eb6f92;
    --gold: #f6c177;
    --rose: #ea9a97;
    --pine: #3e8fb0;
    --foam: #9ccfd8;
    --iris: #c4a7e7;
    --highlight-low: #2a283e;
    --highlight-muted: #44415a;
    --highlight-high: #56526e;

  font-family: Pier Sans, ui-sans-serif, system-ui, -apple-system,
    BlinkMacSystemFont, Segoe UI, Roboto, Helvetica Neue, Arial, Noto Sans,
    sans-serif, "Apple Color Emoji", "Segoe UI Emoji", Segoe UI Symbol,
    "Noto Color Emoji";
  font-weight: initial;

  background-color: var(--base);
}
/*Common style*/
h1 {
  color: var(--rose);
  padding-bottom: 2mm;
  margin-bottom: 12mm;
}
h2 {
  color: var(--rose);
}
h3 {
  color: var(--rose);
}
h4 {
  color: var(--rose);
}
h5 {
  color: var(--rose);
}
h6 {
  color: var(--rose);
}
a {
  color: var(--iris);
}
p {
  font-size: 20pt;
  font-weight: 600;
  color: var(--text);
}
code {
  color: var(--text);
  background-color: var(--highlight-muted);
}
text {
  color: var(--text);
}
ul {
  color: var(--subtle);
}
li {
  color: var(--subtle);
}
img {
  background-color: var(--highlight-low);
}
strong {
  color: var(--text);
  font-weight: inherit;
  font-weight: 800;
}
mjx-container {
  color: var(--text);
}
marp-pre {
  background-color: var(--overlay);
  border-color: var(--highlight-high);
}

/*Code blok*/
.hljs-comment {
  color: var(--muted);
}
.hljs-attr {
  color: var(--foam);
}
.hljs-punctuation {
  color: var(--subtle);
}
.hljs-string {
  color: var(--gold);
}
.hljs-title {
  color: var(--foam);
}
.hljs-keyword {
  color: var(--pine);
}
.hljs-variable {
  color: var(--text);
}
.hljs-literal {
  color: var(--rose);
}
.hljs-type {
  color: var(--love);
}
.hljs-number {
  color: var(--gold);
}
.hljs-built_in {
  color: var(--love);
}
.hljs-params {
  color: var(--iris);
}
.hljs-symbol {
  color: var(--foam);
}
.hljs-meta {
  color: var(--subtle);
}

</style>

<!-- Introduce my self -->
<!-- _class: title-slide -->
<style>
.title-slide {
    display: flex;
    flex-direction: column;
    justify-content: center;
    height: 100%;
}

.title-slide h1 {
    font-size: 40pt;
    margin-bottom: 60px;
    color: var(--pine);
}

.presenter {
    margin-top: 60px;
}

.presenter .name {
    font-size: 28pt;
    color: var(--foam);
    margin-bottom: 10px;
}

.presenter .title {
    font-size: 22pt;
    color: var(--iris);
    margin-bottom: 20px;
}

.social {
    display: flex;
    gap: 15px;
    margin-top: 15px;
}

.social img {
    height: 24px;
    background: transparent;
}
</style>

<div class="color-bar" style="height: 6px; width: 100px; background: linear-gradient(90deg, var(--pine) 70%, var(--gold) 30%);"></div>

## Container Workflows and Orchestration with Sugar

<div class="presenter">
    <div class="name">Sai Sanjay</div>
    <div class="title">GSoC Contributor @ Open Science Labs</div>
    <div class="social">
        <a href="https://github.com/sanjay7178"><img src="https://cdn.jsdelivr.net/npm/simple-icons@v8/icons/github.svg" alt="GitHub"></a>
        <a href="https://linkedin.com/in/sanjay7178"><img src="https://cdn.jsdelivr.net/npm/simple-icons@v8/icons/linkedin.svg" alt="LinkedIn"></a>
        <a href="https://twitter.com/sanjay7178"><img src="https://cdn.jsdelivr.net/npm/simple-icons@v8/icons/twitter.svg" alt="Twitter"></a>
        <span style="color: var(--subtle);">@sanjay7178</span>
    </div>
</div>

---

# Introduction to Kubernetes

---
