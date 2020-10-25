import React from 'react';
import './App.css';
import {GitRepositoryList} from "./components/git-repositories/GitRepositoryList";

function App() {
    return (
        <div className='App'>
            <h2>Git repositories</h2>
            <GitRepositoryList />
        </div>
    )
}

export default App;
