// setup web socket connection
const ws = new WebSocket('ws://localhost:8080/ws');

ws.onopen = function() {
    console.log('WebSocket connection established');
};

ws.onmessage = function(event) {
    const message = JSON.parse(event.data);
    handleWebSocketMessage(message);
};

ws.onerror = function(error) {
    console.log('WebSocket Error:', error);
};

function handleWebSocketMessage(message) {
    if (message.action === 'play') {
        playNote(message.note);
    } else if (message.action === 'stop') {
        stopNote(message.note);
    }
}

// Check if AudioContext is supported
window.AudioContext = window.AudioContext || window.webkitAudioContext;
if (!window.AudioContext) {
    alert("Your browser does not support Web Audio API");
    throw new Error("Web Audio API not supported.");
}

// Initialize AudioContext
const audioCtx = new AudioContext();
// Initialize GainNode
const masterGain = audioCtx.createGain();
masterGain.connect(audioCtx.destination);
masterGain.gain.value = 1; // Start with a default gain of 1

// Frequencies for the C3-B3 notes in the C major scale
const noteFrequencies = {
    'KeyC': 130.81, // C
    'KeyD': 146.83, // D
    'KeyE': 164.81, // E
    'KeyF': 174.61, // F
    'KeyG': 196.00, // G
    'KeyA': 220.00, // A
    'KeyB': 246.94  // B
};

const activeOscillators = {};

function updateMasterGain() {
    // Adjust the gain based on the number of active oscillators
    const activeCount = Object.keys(activeOscillators).length;
    const newGain = activeCount > 0 ? 1 / Math.sqrt(activeCount) : 1;
    masterGain.gain.value = newGain;
}

function playNote (key) {
    if (!noteFrequencies[key]) return;
    const osc = audioCtx.createOscillator();
    osc.frequency.value = noteFrequencies[key];
    osc.type = 'sine';
    osc.connect(masterGain); // Connect to the master gain node instead of directly to destination
    osc.start();
    activeOscillators[key] = osc;
    updateMasterGain();
    updateNoteDisplay(key, true); // Add this line
}

function stopNote (key) {
    if (activeOscillators[key]) {
        activeOscillators[key].stop();
        activeOscillators[key].disconnect();
        delete activeOscillators[key];
        updateMasterGain();
        updateNoteDisplay(key, false); // Add this line
    }
}

window.addEventListener('keydown', function(e) {
    if (e.repeat) return;
    playNote (e.code);
});

window.addEventListener('keyup', function(e) {
    stopNote (e.code);
});

function updateNoteDisplay(note, isActive) {
    const noteDiv = document.getElementById(`note-${note[note.length-1]}`);
    if (noteDiv) {
        if (isActive) {
            // Define colors for each note or use a generic active style
            noteDiv.classList.add('active');
            switch (note) {
                case 'KeyC': noteDiv.style.backgroundColor = 'red'; break;
                case 'KeyD': noteDiv.style.backgroundColor = 'orange'; break;
                case 'KeyE': noteDiv.style.backgroundColor = 'yellow'; break;
                case 'KeyF': noteDiv.style.backgroundColor = 'green'; break;
                case 'KeyG': noteDiv.style.backgroundColor = 'blue'; break;
                case 'KeyA': noteDiv.style.backgroundColor = 'indigo'; break;
                case 'KeyB': noteDiv.style.backgroundColor = 'violet'; break;
            }
        } else {
            noteDiv.classList.remove('active');
            noteDiv.style.backgroundColor = ''; // Reset to default
        }
    }
}