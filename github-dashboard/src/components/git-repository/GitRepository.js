import React, {useState} from "react";
import PropTypes from "prop-types";
import './GitRepository.css'
import githubLogo from '../../assets/icons/github.png'
import {useHttp} from "../../hooks/http";
import {Commit} from "../commit/Commit";

function Contact(props) {
    const language = props.language ? props.language : 'not defined'
    const {data, callApi} = useHttp();
    const [expanded, setExpanded] = useState(false);

    const collapseCommits = () => {
        callApi('http://localhost:8090/commits/' + props.name);
        setExpanded(!expanded);
    }

    return (
        <div className="git-repository-container">
            <div className='git-repository' onClick={() => collapseCommits()}>
                <img className="github" alt="github icon" src={githubLogo}/>
                <span className="name"><b>Repository name: </b>{props.name}</span>
                <span className="language"><b>Written in: </b>{language}</span>
                <p className="description" title={props.description}>{props.description}</p>
            </div>
            <div className={"commits " + expanded}>
                {expanded ? data.map(c => <Commit
                    key={c.node_id}
                    name={c.commit.author.name}
                    html_url={c.html_url}
                    description={c.commit.message}/>) : ""}
            </div>
        </div>
    );
}

Contact.propTypes = {
    name: PropTypes.string.isRequired,
};

export default Contact;