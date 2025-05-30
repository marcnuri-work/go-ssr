import * as React from 'react';
import {useState} from 'react';
import {renderToString} from 'react-dom/server';
import {createRoot} from 'react-dom/client';

const App = () => {
  const [count, setCount] = useState(0);
  const increment = () => setCount(count + 1);
  return (
    <div>
      <h1>Greetings professor Falken!</h1>
      <button onClick={increment}>Hey {count} times!</button>
    </div>
  );
};

const renderedString = renderToString(
  <App />
);

// goSsr is a global object provided by the Go SSR framework
goSsr.render(renderedString);

// Client-side rendering
debugger
if (typeof document !== 'undefined') {
  const container = document.getElementById('root');
  if (container) {
    const root = createRoot(container);
    root.render(<App />);
  }
}
