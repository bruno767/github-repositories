import React from 'react';
import {Loading} from '../loading/loading';
import GitRepository from "../git-repository/GitRepository";
import {useHttp} from "../../hooks/http";
import './GitRepositoryList.css'

export function GitRepositoryList() {

    const [loading, gitRepositories] = useHttp('http://localhost:8090/repositories', [])

    if (loading) return <Loading/>

    return (
        <div className='gitRepositories'>
            {gitRepositories.map(g => <GitRepository
                key={g.id}
                name={g.name}
                html_url={g.html_url}
                description={g.description}
                language={g.language}/>)}
        </div>
    )

}