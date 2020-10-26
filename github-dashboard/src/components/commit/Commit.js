import React from 'react';
import './Commit.css'

export function Commit(props) {
    return (
        <div className='commit'>
            <span className="name"><b>Author name: </b>{props.name}</span>
            <p className="description" title={props.description}>{props.description}</p>
        </div>
    )

}