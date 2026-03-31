const board = document.getElementById("board");
const keyboardDiv = document.getElementById("keyboard");
const messageDiv = document.getElementById("message");

let word = "";
const maxRows = 5;
let currentRow = 0;
let guesses = JSON.parse(localStorage.getItem("quortle-guesses") || "[]");

// QWERTY layout
const keys = [
  "QWERTYUIOP".split(""),
  "ASDFGHJKL".split(""),
  ["Enter", "Z", "X", "C", "V", "B", "N", "M", "⌫"]
];

let validWords = new Set();

async function loadValidWords() {
  try {
    const res = await fetch("../words.txt");
    const text = await res.text();
    const words = text.split("\n").map(w => w.trim().toLowerCase()).filter(w => w.length === 4);
    validWords = new Set(words);
    console.log(`Loaded ${validWords.size} valid words.`);
  } catch (err) {
    console.error("Failed to load words.txt", err);
  }
}

// Initialize board
function initBoard() {
  board.innerHTML = "";
  for (let r = 0; r < maxRows; r++) {
    for (let c = 0; c < 4; c++) {
      const cell = document.createElement("div");
      cell.classList.add("cell");
      cell.dataset.row = r;
      cell.dataset.col = c;
      board.appendChild(cell);
    }
  }
}

// Initialize keyboard
function initKeyboard() {
  keyboardDiv.innerHTML = "";
  keys.flat().forEach(k => {
    const key = document.createElement("button");
    key.textContent = k;
    key.classList.add("key");
    key.addEventListener("click", (e) => {
      handleKey(k);
      e.target.blur(); //  THIS FIXES IT
    }); keyboardDiv.appendChild(key);
  });
}

// Fetch today's word
async function fetchWord() {
  try {
    const res = await fetch("/word/random");
    const data = await res.json();
    word = data.word.toLowerCase();
    console.log("For testing purposes a daily word is printed here: ", word);
  } catch {
    messageDiv.textContent = "Error fetching word";
  }
}

// Handle key input
function handleKey(k) {
  if (currentRow >= maxRows) return;

  let guess = getCurrentGuess();
  if (k === "backspace" || k === "⌫") {
    guess = guess.slice(0, -1);
    setCurrentGuess(guess);
  } else if (k === "Enter") {
    if (guess.length < 4) {
      alert("Enter a 4-letter word");
      return;
    }
    if (!validWords.has(guess.toLowerCase())) {
      // Word not valid
      messageDiv.textContent = "Not a valid word";
      shakeRow(currentRow); // Visual feedback
      return;
    }
    // Word is valid
    checkGuess(guess);
  } else if (guess.length < 4 && k.length === 1) {
    guess += k.toLowerCase();
    setCurrentGuess(guess);
  }
}

// Get current guess
function getCurrentGuess() {
  let guess = "";
  for (let c = 0; c < 4; c++) {
    const cell = document.querySelector(`.cell[data-row="${currentRow}"][data-col="${c}"]`);
    guess += cell.textContent.toLowerCase();
  }
  return guess;
}

// Set current guess
function setCurrentGuess(guess) {
  for (let c = 0; c < 4; c++) {
    const cell = document.querySelector(`.cell[data-row="${currentRow}"][data-col="${c}"]`);
    cell.textContent = guess[c] || "";
  }
}

// Check guess and update colors
function checkGuess(guess) {
  guess = guess.toLowerCase();
  const letters = guess.split("");
  const wordLetters = word.split("");
  const keyElements = Array.from(document.querySelectorAll(".key"));

  letters.forEach((l, i) => {
    const cell = document.querySelector(`.cell[data-row="${currentRow}"][data-col="${i}"]`);
    if (l === wordLetters[i]) {
      cell.classList.add("correct");
      const keyBtn = keyElements.find(k => k.textContent.toLowerCase() === l);
      if (keyBtn) keyBtn.classList.add("correct");
    } else if (wordLetters.includes(l)) {
      cell.classList.add("present");
      const keyBtn = keyElements.find(k => k.textContent.toLowerCase() === l);
      if (keyBtn && !keyBtn.classList.contains("correct")) keyBtn.classList.add("present");
    } else {
      cell.classList.add("absent");
      const keyBtn = Array.from(document.querySelectorAll(".key"))
        .find(k => k.textContent.toLowerCase() === l);

      if (keyBtn && !keyBtn.classList.contains("correct") && !keyBtn.classList.contains("present")) keyBtn.classList.add("absent");
    }
    cell.style.transform = "scale(1.1)";
    setTimeout(() => { cell.style.transform = "scale(1)"; }, 200);
  });

  guesses.push(guess);
  localStorage.setItem("quortle-guesses", JSON.stringify(guesses));

  if (guess === word) {
    messageDiv.textContent = "Correct! The word is " + word.toUpperCase();
    currentRow = maxRows;
    localStorage.removeItem("quortle-guesses");
  } else {
    currentRow++;
    if (currentRow >= maxRows) {
      messageDiv.textContent = "Game over! \n The word was " + word.toUpperCase();
      localStorage.removeItem("quortle-guesses");
    }
  }
}

// Handle physical keyboard input
document.addEventListener("keydown", (e) => {
  if (currentRow >= maxRows) return;

  let key = e.key;

  if (key === "Enter") {
    handleKey("Enter");
  } else if (key === "⌫" || key === "Backspace") {
    handleKey("⌫");
  } else if (/^[a-zA-Z]$/.test(key)) {  // only letters
    handleKey(key.toUpperCase());
  }
});

function shakeRow(row) {
  for (let c = 0; c < 4; c++) {
    const cell = document.querySelector(`.cell[data-row="${row}"][data-col="${c}"]`);
    cell.classList.add("shake");
    setTimeout(() => cell.classList.remove("shake"), 300);
  }
}

// Restore previous guesses
function restoreGuesses() {
  guesses.forEach(g => {
    setCurrentGuess(g);
    checkGuess(g);
  });
  currentRow = guesses.length;
}

// Initialize
window.onload = async () => {
  await loadValidWords();
  await fetchWord();
  initBoard();
  initKeyboard();
  restoreGuesses();
};

const footerYear = document.querySelectorAll(".year");
footerYear.forEach(copyright => {
  copyright.innerHTML = new Date().getFullYear();
});