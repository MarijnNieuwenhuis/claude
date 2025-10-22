# Chota CSS Framework - Complete Reference

**Version:** Latest (3kb minified + gzipped)
**Official Site:** https://jenil.github.io/chota/
**Philosophy:** A micro CSS framework requiring no preprocessor, just plug-n-play

The name "Chota" comes from Hindi, meaning "small."

---

## Table of Contents

1. [Installation](#installation)
2. [Customization & Theming](#customization--theming)
3. [Grid System](#grid-system)
4. [Typography](#typography)
5. [Buttons](#buttons)
6. [Forms](#forms)
7. [Navigation](#navigation)
8. [Cards](#cards)
9. [Tags](#tags)
10. [Tables](#tables)
11. [Details & Dropdowns](#details--dropdowns)
12. [Utilities](#utilities)
13. [Icons](#icons)

---

## Installation

### CDN (Recommended for Quick Start)

```html
<link rel="stylesheet" href="https://unpkg.com/chota@latest">
```

### NPM

```bash
npm install chota
```

Then import in your project:
```javascript
import 'chota';
```

---

## Customization & Theming

Chota uses CSS variables for easy theming. Override by adding custom values to `:root` **after** importing the framework.

### Core CSS Variables

```css
:root {
  /* Colors */
  --bg-color: #ffffff;
  --bg-secondary-color: #f3f3f6;
  --color-primary: #14854F;
  --color-lightGrey: #d2d6dd;
  --color-grey: #747681;
  --color-darkGrey: #3f4144;
  --color-error: #d43939;
  --color-success: #28bd14;

  /* Grid */
  --grid-maxWidth: 120rem;
  --grid-gutter: 2rem;

  /* Typography */
  --font-size: 1.6rem;
  --font-family-sans: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif;
  --font-family-mono: monaco, "Consolas", "Lucida Console", monospace;
}
```

### Dark Mode

Chota includes built-in dark mode support. Add the `dark` class to `<body>` to enable:

```html
<body class="dark">
  <!-- Your content -->
</body>
```

The framework automatically detects device dark mode preferences and adjusts colors accordingly.

### Custom Theme Example

```css
:root {
  --color-primary: #0066cc;
  --color-error: #ff4444;
  --grid-maxWidth: 100rem;
  --font-size: 1.4rem;
}
```

---

## Grid System

Chota provides a flexible, responsive 12-column grid system.

### Basic Grid

Use `.row` containers with `.col` children for equal-width columns:

```html
<div class="row">
  <div class="col">Column 1</div>
  <div class="col">Column 2</div>
  <div class="col">Column 3</div>
</div>
```

### Sized Columns

Specify column widths using `.col-1` through `.col-12`:

```html
<div class="row">
  <div class="col-3">25% width</div>
  <div class="col-6">50% width</div>
  <div class="col-3">25% width</div>
</div>
```

### Responsive Columns

Chota provides three responsive breakpoints:

- **Default:** All sizes ≥600px (mobile-first)
- **Medium (md):** ≥900px
- **Large (lg):** ≥1200px

**Below 600px, all columns default to 100% width.**

```html
<div class="row">
  <div class="col-12 col-6-md col-4-lg">
    <!-- Full width on mobile, half on tablet, third on desktop -->
  </div>
  <div class="col-12 col-6-md col-8-lg">
    <!-- Full width on mobile, half on tablet, 2/3 on desktop -->
  </div>
</div>
```

### Combining Flexible & Sized Columns

You can mix `.col` and sized columns:

```html
<div class="row">
  <div class="col-4">Fixed 4 columns</div>
  <div class="col">Fills remaining space</div>
</div>
```

### Reverse Direction

Reverse column order with `.reverse`:

```html
<div class="row reverse">
  <div class="col">Appears second visually</div>
  <div class="col">Appears first visually</div>
</div>
```

### Force Line Breaks

Use `.is-full-width` to force a column to take the full row width:

```html
<div class="row">
  <div class="col-6">First column</div>
  <div class="col-6 is-full-width">Breaks to new line</div>
  <div class="col-6">Third column on another line</div>
</div>
```

---

## Typography

Chota provides clean, readable typography out of the box.

### Headings

```html
<h1>Heading 1</h1>
<h2>Heading 2</h2>
<h3>Heading 3</h3>
<h4>Heading 4</h4>
<h5>Heading 5</h5>
<h6>Heading 6</h6>
```

### Paragraphs & Text Elements

```html
<p>Standard paragraph text with automatic line-height and spacing.</p>
<strong>Bold text</strong>
<em>Italic text</em>
<u>Underlined text</u>
<del>Strikethrough text</del>
<code>Inline code</code>
```

### Code Blocks

```html
<pre><code>
function hello() {
  console.log("Formatted code block");
}
</code></pre>
```

### Blockquotes

```html
<blockquote>
  This is a quotation with automatic styling and left border.
</blockquote>
```

### Lists

```html
<ul>
  <li>Unordered list item 1</li>
  <li>Unordered list item 2</li>
</ul>

<ol>
  <li>Ordered list item 1</li>
  <li>Ordered list item 2</li>
</ol>
```

### Horizontal Rule

```html
<hr>
```

---

## Buttons

Chota provides various button styles and states.

### Button Types

```html
<!-- Default button -->
<button>Default Button</button>

<!-- Primary button -->
<button class="primary">Primary Button</button>

<!-- Secondary button -->
<button class="secondary">Secondary Button</button>

<!-- Dark button -->
<button class="dark">Dark Button</button>

<!-- Error button -->
<button class="error">Error Button</button>

<!-- Success button -->
<button class="success">Success Button</button>
```

### Outline Buttons

Add `outline` class for outlined variants:

```html
<button class="outline">Default Outline</button>
<button class="primary outline">Primary Outline</button>
<button class="secondary outline">Secondary Outline</button>
<button class="dark outline">Dark Outline</button>
```

### Clear Buttons

Minimal styling with no background or border:

```html
<button class="clear">Clear Button</button>
```

### Icon Buttons

Add icons using the `icon` class, or create icon-only buttons:

```html
<!-- Button with icon -->
<button class="icon">
  <img src="icon.svg" alt="">
  Button Text
</button>

<!-- Icon-only button -->
<button class="icon-only">
  <img src="icon.svg" alt="">
</button>
```

### Button States

```html
<!-- Disabled button -->
<button disabled>Disabled Button</button>

<!-- Loading state (add via JavaScript) -->
<button class="loading">Loading...</button>
```

### Link Buttons

Style links as buttons:

```html
<a class="button" href="#">Link as Button</a>
<a class="button primary" href="#">Primary Link Button</a>
```

---

## Forms

Chota styles form elements with consistent spacing and visual feedback.

### Input Fields

```html
<label for="name">Name</label>
<input type="text" id="name" placeholder="Enter your name">

<label for="email">Email</label>
<input type="email" id="email" placeholder="your@email.com">

<label for="password">Password</label>
<input type="password" id="password">
```

### Textarea

```html
<label for="message">Message</label>
<textarea id="message" rows="4" placeholder="Your message here..."></textarea>
```

### Select Dropdowns

```html
<label for="country">Country</label>
<select id="country">
  <option>Select a country</option>
  <option value="us">United States</option>
  <option value="uk">United Kingdom</option>
  <option value="ca">Canada</option>
</select>
```

### Checkboxes

```html
<label>
  <input type="checkbox">
  I agree to the terms and conditions
</label>

<label>
  <input type="checkbox" checked>
  Pre-checked option
</label>
```

### Radio Buttons

```html
<fieldset>
  <legend>Choose an option</legend>

  <label>
    <input type="radio" name="option" value="1">
    Option 1
  </label>

  <label>
    <input type="radio" name="option" value="2" checked>
    Option 2
  </label>

  <label>
    <input type="radio" name="option" value="3">
    Option 3
  </label>
</fieldset>
```

### Form Validation States

```html
<!-- Error state -->
<input type="text" class="error" placeholder="Invalid input">
<span class="text-error">This field is required</span>

<!-- Success state -->
<input type="text" class="success" placeholder="Valid input">
<span class="text-success">Looks good!</span>
```

### Input Groups

```html
<div class="grouped">
  <input type="text" placeholder="Username">
  <button>Submit</button>
</div>
```

### Full Form Example

```html
<form>
  <div class="row">
    <div class="col">
      <label for="firstname">First Name</label>
      <input type="text" id="firstname" required>
    </div>
    <div class="col">
      <label for="lastname">Last Name</label>
      <input type="text" id="lastname" required>
    </div>
  </div>

  <label for="email">Email</label>
  <input type="email" id="email" required>

  <label for="message">Message</label>
  <textarea id="message" rows="4"></textarea>

  <label>
    <input type="checkbox" required>
    I agree to receive updates
  </label>

  <button type="submit" class="primary">Submit</button>
</form>
```

---

## Navigation

Chota provides flexible navigation components.

### Basic Navigation

Use `.nav` container with three optional sections:

```html
<nav class="nav">
  <div class="nav-left">
    <a class="brand" href="#">Brand</a>
  </div>

  <div class="nav-center">
    <a href="#">About</a>
    <a href="#" class="active">Services</a>
    <a href="#">Contact</a>
  </div>

  <div class="nav-right">
    <button>Sign In</button>
  </div>
</nav>
```

### Navigation with Tabs

```html
<nav class="nav">
  <div class="nav-left">
    <a class="brand" href="#">
      <img src="logo.svg" alt="Logo">
    </a>
  </div>

  <div class="nav-center">
    <div class="tabs">
      <a href="#" class="active">Home</a>
      <a href="#">Products</a>
      <a href="#">About</a>
      <a href="#">Contact</a>
    </div>
  </div>

  <div class="nav-right">
    <button class="primary">Get Started</button>
  </div>
</nav>
```

### Standalone Tabs

Tabs can be used independently outside navigation:

```html
<div class="tabs">
  <a href="#" class="active">Tab 1</a>
  <a href="#">Tab 2</a>
  <a href="#">Tab 3</a>
</div>
```

### Full-Width Tabs

Add `is-full` for full-width tab styling:

```html
<div class="tabs is-full">
  <a href="#" class="active">Overview</a>
  <a href="#">Specifications</a>
  <a href="#">Reviews</a>
</div>
```

### Active State

Use `.active` class to highlight the current page/tab:

```html
<a href="#" class="active">Current Page</a>
```

---

## Cards

Cards provide a container with shadow, border, radius, and padding.

### Basic Card

```html
<div class="card">
  <p>This is a simple card with default styling.</p>
</div>
```

### Card with Header

```html
<div class="card">
  <header>
    <h3>Card Title</h3>
  </header>
  <p>Card content goes here with automatic spacing.</p>
</div>
```

### Card with Footer

```html
<div class="card">
  <header>
    <h3>Product Name</h3>
  </header>
  <p>Product description and details.</p>
  <footer>
    <button class="primary">Buy Now</button>
    <button class="outline">Learn More</button>
  </footer>
</div>
```

### Complete Card Example

```html
<div class="row">
  <div class="col-4">
    <div class="card">
      <header>
        <h4>Basic Plan</h4>
      </header>
      <p>Perfect for individuals</p>
      <ul>
        <li>5 GB Storage</li>
        <li>Email Support</li>
        <li>Basic Features</li>
      </ul>
      <footer>
        <button class="primary">Choose Plan</button>
      </footer>
    </div>
  </div>

  <div class="col-4">
    <div class="card">
      <header>
        <h4>Pro Plan</h4>
      </header>
      <p>For growing teams</p>
      <ul>
        <li>50 GB Storage</li>
        <li>Priority Support</li>
        <li>Advanced Features</li>
      </ul>
      <footer>
        <button class="primary">Choose Plan</button>
      </footer>
    </div>
  </div>

  <div class="col-4">
    <div class="card">
      <header>
        <h4>Enterprise</h4>
      </header>
      <p>For large organizations</p>
      <ul>
        <li>Unlimited Storage</li>
        <li>24/7 Support</li>
        <li>Custom Solutions</li>
      </ul>
      <footer>
        <button class="primary">Contact Us</button>
      </footer>
    </div>
  </div>
</div>
```

---

## Tags

Labels/tags in three sizes for categorization and status indicators.

### Tag Sizes

```html
<!-- Small tag -->
<span class="tag is-small">Small Tag</span>

<!-- Default tag -->
<span class="tag">Default Tag</span>

<!-- Large tag -->
<span class="tag is-large">Large Tag</span>
```

### Colored Tags

Combine with color utility classes:

```html
<span class="tag bg-primary">Primary</span>
<span class="tag bg-error">Error</span>
<span class="tag bg-success">Success</span>
<span class="tag bg-dark text-white">Dark</span>
```

### Usage Example

```html
<h3>
  Article Title
  <span class="tag is-small bg-success">New</span>
</h3>

<p>
  Tags:
  <span class="tag">JavaScript</span>
  <span class="tag">CSS</span>
  <span class="tag">HTML</span>
</p>
```

---

## Tables

Chota styles tables with proper spacing and borders.

### Basic Table

```html
<table>
  <thead>
    <tr>
      <th>Name</th>
      <th>Email</th>
      <th>Role</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>John Doe</td>
      <td>john@example.com</td>
      <td>Admin</td>
    </tr>
    <tr>
      <td>Jane Smith</td>
      <td>jane@example.com</td>
      <td>Editor</td>
    </tr>
    <tr>
      <td>Bob Johnson</td>
      <td>bob@example.com</td>
      <td>Viewer</td>
    </tr>
  </tbody>
</table>
```

### Striped Table

Add `striped` class for alternating row colors:

```html
<table class="striped">
  <!-- table content -->
</table>
```

### Responsive Table

Wrap table in a container with `is-full-width` for horizontal scrolling on small screens:

```html
<div class="is-full-width">
  <table>
    <!-- many columns -->
  </table>
</div>
```

---

## Details & Dropdowns

Native `<details>` elements with enhanced styling.

### Basic Details

```html
<details>
  <summary>Click to expand</summary>
  <p>Hidden content revealed when opened.</p>
</details>
```

### Open by Default

```html
<details open>
  <summary>Already expanded</summary>
  <p>This content is visible by default.</p>
</details>
```

### Dropdown Menu

Combine with `.dropdown`, `.button`, and `.card`:

```html
<details class="dropdown">
  <summary class="button">
    Menu
  </summary>
  <div class="card">
    <a href="#">Option 1</a>
    <a href="#">Option 2</a>
    <a href="#">Option 3</a>
  </div>
</details>
```

### Styled Dropdown Example

```html
<div class="nav-right">
  <details class="dropdown">
    <summary class="button outline">
      Account
    </summary>
    <div class="card">
      <a href="#">Profile</a>
      <a href="#">Settings</a>
      <hr>
      <a href="#" class="text-error">Logout</a>
    </div>
  </details>
</div>
```

---

## Utilities

Chota provides extensive utility classes for common styling needs.

### Text Colors

```html
<p class="text-primary">Primary colored text</p>
<p class="text-light">Light colored text</p>
<p class="text-white">White colored text</p>
<p class="text-dark">Dark colored text</p>
<p class="text-grey">Grey colored text</p>
<p class="text-error">Error colored text</p>
<p class="text-success">Success colored text</p>
```

### Text Alignment

```html
<p class="text-left">Left aligned text</p>
<p class="text-center">Center aligned text</p>
<p class="text-right">Right aligned text</p>
```

### Text Transform

```html
<p class="text-uppercase">UPPERCASE TEXT</p>
<p class="text-lowercase">lowercase text</p>
<p class="text-capitalize">Capitalized Text</p>
```

### Background Colors

```html
<div class="bg-primary">Primary background</div>
<div class="bg-dark text-white">Dark background</div>
<div class="bg-grey">Grey background</div>
<div class="bg-error text-white">Error background</div>
<div class="bg-success text-white">Success background</div>
```

### Border Colors

```html
<div class="bd-primary">Primary border</div>
<div class="bd-dark">Dark border</div>
<div class="bd-grey">Grey border</div>
<div class="bd-error">Error border</div>
<div class="bd-success">Success border</div>
```

### Positioning

```html
<!-- Float elements -->
<div class="pull-right">Floated right</div>
<div class="pull-left">Floated left</div>

<!-- Fixed positioning -->
<div class="is-fixed">Fixed position element</div>

<!-- Alignment -->
<div class="is-center">Centered element</div>
<div class="is-right">Right aligned element</div>
<div class="is-left">Left aligned element</div>
```

### Layout Utilities

```html
<!-- Full width/screen -->
<div class="is-full-width">Full width element</div>
<div class="is-full-screen">Full screen element</div>

<!-- Vertical/horizontal alignment -->
<div class="is-vertical-align">Vertically aligned</div>
<div class="is-horizontal-align">Horizontally aligned</div>
```

### Spacing

```html
<!-- Remove padding -->
<div class="is-paddingless">No padding</div>

<!-- Remove margin -->
<div class="is-marginless">No margin</div>
```

### Visibility

```html
<!-- Hide element completely -->
<div class="is-hidden">Hidden element</div>

<!-- Responsive visibility -->
<div class="hide-xs">Hidden on extra small screens</div>
<div class="hide-sm">Hidden on small screens</div>
<div class="hide-md">Hidden on medium screens</div>
<div class="hide-lg">Hidden on large screens</div>
<div class="hide-pr">Hidden in print</div>
```

### Other Utilities

```html
<!-- Rounded corners -->
<div class="is-rounded">Rounded element</div>

<!-- Clearfix for floated children -->
<div class="clearfix">
  <div class="pull-left">Float left</div>
  <div class="pull-right">Float right</div>
</div>
```

---

## Icons

Chota doesn't include icons but integrates easily with icon libraries.

### Using Icongram

Icongram provides quick icon integration:

```html
<img src="https://icongr.am/feather/home.svg?size=24&color=currentColor" alt="Home">
```

### Font Awesome

```html
<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.0.0/css/all.min.css">

<i class="fas fa-home"></i>
<i class="fas fa-user"></i>
<i class="fas fa-cog"></i>
```

### Icon in Button

```html
<button class="icon">
  <img src="icon.svg" alt="">
  Button with Icon
</button>

<button class="icon-only primary">
  <img src="icon.svg" alt="">
</button>
```

---

## Responsive Breakpoints Summary

Chota uses three main breakpoints:

| Breakpoint | Min Width | Class Suffix | Usage |
|------------|-----------|--------------|-------|
| Mobile (default) | < 600px | (none) | Mobile-first base styles |
| Tablet | ≥ 600px | (default) | Small tablets and up |
| Medium | ≥ 900px | `-md` | Tablets landscape and up |
| Large | ≥ 1200px | `-lg` | Desktop and up |

**Example:**
```html
<div class="col-12 col-6-md col-4-lg">
  <!-- Full width on mobile -->
  <!-- Half width on medium screens -->
  <!-- One-third width on large screens -->
</div>
```

---

## Best Practices

### 1. Mobile-First Approach
Start with mobile styles, then enhance for larger screens:

```html
<div class="col-12 col-6-md col-4-lg">
  <!-- Progressively enhance from mobile up -->
</div>
```

### 2. Semantic HTML
Use appropriate HTML elements with Chota's styling:

```html
<!-- Good -->
<button class="primary">Submit</button>

<!-- Avoid -->
<div class="button primary" onclick="submit()">Submit</div>
```

### 3. Leverage CSS Variables
Customize the theme by overriding variables:

```css
:root {
  --color-primary: #007bff;
  --grid-maxWidth: 100rem;
}
```

### 4. Combine Classes
Mix and match utilities for precise control:

```html
<button class="primary outline icon">
  <img src="icon.svg" alt="">
  Combined Styles
</button>
```

### 5. Use Cards for Grouping
Cards provide visual hierarchy and spacing:

```html
<div class="card">
  <header><h3>Section Title</h3></header>
  <p>Related content grouped together</p>
</div>
```

---

## Common Patterns

### Hero Section

```html
<div class="row is-full-screen is-center">
  <div class="col-8 text-center">
    <h1>Welcome to Our Site</h1>
    <p>Build fast, responsive websites with Chota</p>
    <button class="primary">Get Started</button>
    <button class="outline">Learn More</button>
  </div>
</div>
```

### Feature Grid

```html
<div class="row">
  <div class="col-4">
    <div class="card text-center">
      <h3>Fast</h3>
      <p>Only 3kb gzipped</p>
    </div>
  </div>
  <div class="col-4">
    <div class="card text-center">
      <h3>Simple</h3>
      <p>No build process needed</p>
    </div>
  </div>
  <div class="col-4">
    <div class="card text-center">
      <h3>Flexible</h3>
      <p>Easy to customize</p>
    </div>
  </div>
</div>
```

### Form with Validation

```html
<form>
  <label for="email">Email</label>
  <input type="email" id="email" class="error">
  <span class="text-error">Please enter a valid email</span>

  <label for="password">Password</label>
  <input type="password" id="password" class="success">
  <span class="text-success">Strong password!</span>

  <button type="submit" class="primary">Sign In</button>
</form>
```

### Navigation with Dropdown

```html
<nav class="nav">
  <div class="nav-left">
    <a class="brand" href="#">Logo</a>
  </div>

  <div class="nav-center">
    <a href="#" class="active">Home</a>
    <a href="#">About</a>
    <details class="dropdown">
      <summary>Products</summary>
      <div class="card">
        <a href="#">Product A</a>
        <a href="#">Product B</a>
        <a href="#">Product C</a>
      </div>
    </details>
  </div>

  <div class="nav-right">
    <button class="primary">Contact</button>
  </div>
</nav>
```

---

## Browser Support

Chota supports all modern browsers:
- Chrome (latest)
- Firefox (latest)
- Safari (latest)
- Edge (latest)

IE11 is not officially supported due to CSS variable usage.

---

## File Size

- **Uncompressed:** ~12kb
- **Minified:** ~8kb
- **Minified + Gzipped:** ~3kb

---

## License

MIT License - Free for personal and commercial use.

---

## Resources

- **Official Website:** https://jenil.github.io/chota/
- **GitHub Repository:** https://github.com/jenil/chota
- **NPM Package:** https://www.npmjs.com/package/chota
- **CDN:** https://unpkg.com/chota@latest

---

## Quick Reference Card

### Most Common Classes

| Purpose | Classes |
|---------|---------|
| Grid | `.row`, `.col`, `.col-1` to `.col-12` |
| Responsive | `.col-N-md`, `.col-N-lg` |
| Buttons | `.primary`, `.secondary`, `.dark`, `.outline`, `.clear` |
| Text Color | `.text-primary`, `.text-error`, `.text-success` |
| Background | `.bg-primary`, `.bg-dark`, `.bg-grey` |
| Alignment | `.text-center`, `.text-left`, `.text-right` |
| Visibility | `.is-hidden`, `.hide-xs`, `.hide-sm`, `.hide-md`, `.hide-lg` |
| Spacing | `.is-paddingless`, `.is-marginless` |
| Layout | `.is-full-width`, `.is-center`, `.pull-right`, `.pull-left` |
| Components | `.nav`, `.card`, `.tag`, `.tabs`, `.dropdown` |

---

*This documentation is based on Chota CSS Framework. For the latest updates, visit the official website.*
