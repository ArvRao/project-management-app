import { Component, ReactNode } from "react"

class About extends Component {
    state = {
        name: "Arvind"
    }

    render() {
        return (
        <>
        <h1>About page</h1>
        <button>Whats it about?</button>
        </>
        )
    }
}

export default About