class Quiz extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            correct: null
        };
    }

    answer(correct) {
        this.setState({
            correct
        });
    }

    render() {
        conosle.log("Is this thing on.");
        let question = null;
        if (this.state.correct === null) {
            question = <div>
                <button onClick={() => this.answer(false)}>a) A graph database query language</button>
                <button onClick={() => this.answer(true)}>b) An API query language</button>
                <button onClick={() => this.answer(false)}>c) A graph drawing API</button>
            </div>;
        }
        let answer = null;
        if (this.state.correct === true) {
            answer = <div className="correct">Correct! It is an API query language</div>;
        } else if (this.state.correct === false) {
            answer = <div className="incorrect">Nope! It's actually an API query language</div>;
        }

        return <div className="quiz">
            <p>What is GraphQL?</p>
            {question}
            {answer}
        </div>;
    }
}
ReactDOM.render(
    <Quiz />,
    document.getElementById('quiz')
);
/* Taken from https://meda.io/embed-react-into-an-html-web-page/ temporarily for the sake of learning */
