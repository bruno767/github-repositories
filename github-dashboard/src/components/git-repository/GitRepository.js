import React from "react";
import PropTypes from "prop-types";
import './GitRepository.css'
import githubLogo from '../../assets/icons/github.png'

function Contact(props) {
    const language = props.language ? props.language : 'not defined'
    return (
        <a className='git-repository' href={props.html_url}>
            <img className="github" alt="github icon" src={githubLogo}/>
            <span className="name"><b>Repository name: </b>{props.name}</span>
            <span className="language"><b>Written in: </b>{language}</span>
            <p className="description" title={props.description}>{props.description}</p>
        </a>
    );
}

Contact.propTypes = {
    name: PropTypes.string.isRequired,
};

export default Contact;