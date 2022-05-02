import {useEffect, useState} from 'react'
import sendRequest from './sendRequest'

function App() {
  const [state, setState] = useState({
    firstOperand: 0,
    secondOperand: 0,
    result: 0
  })

  const [history, setHistory] = useState([])

  useEffect(() => {
    updateHistory();
  }, [])

  const calculate = (operationSymbol) => {
    sendRequest('POST', 'math.calculate', {
      first_operand: state.firstOperand,
      second_operand: state.secondOperand,
      operation_symbol: operationSymbol
    }).then(({data}) => {
      setState(prevState => {
        updateHistory();

        return {...prevState, result: data.result}
      })
    });
  }

  const updateHistory = () => {
    sendRequest('GET', 'history.get')
      .then(({data}) => {
        const updatedHistory = [];

        data.forEach(element => {
          updatedHistory.push({
            id: element.id,
            firstOperand: element.first_operand,
            secondOperand: element.second_operand,
            operationSymbol: element.operation_symbol,
            result: element.result,
            createdAt: element.created_at
          })
        });

        setHistory(updatedHistory)
      });
  }

  const clearHistory = () => {
    sendRequest('DELETE', 'history.clear')
      .then(() => {
        updateHistory();
      })
  }

  const historyEntryClickHandler = (entryId) => {
    setState(prevState => {
      const historyEntry = history.find(obj => (obj.id === entryId));

      return {
        firstOperand: historyEntry.firstOperand,
        secondOperand: historyEntry.secondOperand,
        result: historyEntry.result
      }
    })
  }

  return (
    <div className="wrapper">
      <div className="input-fields">
        <input
          type="number"
          id="first-input"
          value={state.firstOperand}
          onChange={e => setState(prevState => {
            return {...prevState, firstOperand: +e.target.value}
          })}
        />
        <button className="button" onClick={() => calculate('+')}>+</button>
        <button className="button" onClick={() => calculate('*')}>&times;</button>
        <input
          type="number"
          id="second-input"
          value={state.secondOperand}
          onChange={e => setState(prevState => {
            return {...prevState, secondOperand: +e.target.value}
          })}
        />
      </div>

      <p className="result">Результат: {state.result}</p>

      <hr/>

      <div className="history">
        <p>История:</p>
        <div className="equations">
          {history.map(entry => {
            const options = {
              year: 'numeric',
              month: 'long',
              day: 'numeric',
              hour: 'numeric',
              minute: 'numeric',
              second: 'numeric'
            }

            const stringDate = new Date(entry.createdAt).toLocaleString('ru', options);

            return (<p
              className="history-entry"
              key={entry.id}
              onClick={() => historyEntryClickHandler(entry.id)}
            >
              [{stringDate}] {`${entry.firstOperand} ${entry.operationSymbol} ${entry.secondOperand} = ${entry.result}`}
            </p>)
          })}
        </div>
        <button className="button" onClick={clearHistory}>Очистить</button>
      </div>
    </div>
  );
}

export default App;
