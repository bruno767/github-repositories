import React, {useEffect} from 'react';
import {Loading} from '../loading/loading';
import GitRepository from "../git-repository/GitRepository";
import {useHttp} from "../../hooks/http";
import './GitRepositoryList.css'

export function GitRepositoryList() {

    const {loading, data, callApi} = useHttp();

    useEffect(() => {
        callApi('http://localhost:8090/repositories');
    },[]);

    if (loading) return <Loading/>

    return (
        <div className='gitRepositories'>
            <div className="column">
                {data.filter((element, index, array) => index % 2 === 0).map(g => <GitRepository
                    key={g.id}
                    name={g.name}
                    html_url={g.html_url}
                    description={g.description}
                    language={g.language}/>)}
            </div>
            <div className="column">
                {data.filter((element, index, array) => index % 2 !== 0).map(g => <GitRepository
                    key={g.id}
                    name={g.name}
                    html_url={g.html_url}
                    description={g.description}
                    language={g.language}/>)}
            </div>
        </div>
    )

}