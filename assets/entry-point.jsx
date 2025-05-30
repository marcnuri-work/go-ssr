import * as React from 'react';
import {renderToString} from 'react-dom/server';

const App = () => (
  <div>
    <h1>Greetings professor Falken!</h1>
  </div>
);

const renderedString = renderToString(
  <App />
);
// goSsr is a global object provided by the Go SSR framework
goSsr.render(renderedString);
