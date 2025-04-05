// TinyBet Prediction Market Application

// Contract ABIs (these will be filled from compiled contracts)
const TinyOracleABI = [
    // ABI will be replaced with actual compiled contract ABI
];

const TinyBetABI = [
    // ABI will be replaced with actual compiled contract ABI
];

// App class to handle all functionality
class PredictionMarketApp {
    constructor() {
        this.web3 = null;
        this.accounts = [];
        this.currentAccount = null;
        this.apiBaseUrl = 'http://localhost:8080/api/v1';
        
        // Store deployed contracts
        this.deployedOracles = [];
        this.deployedBets = [];
        
        // Initialize the app
        this.init();
    }
    
    async init() {
        // Set up page navigation
        this.setupNavigation();
        
        // Setup wallet connection
        this.setupWalletConnection();
        
        // Setup form submissions
        this.setupForms();
        
        // Check for MetaMask
        if (window.ethereum) {
            // Check if already connected
            try {
                const accounts = await window.ethereum.request({ method: 'eth_accounts' });
                if (accounts.length > 0) {
                    this.handleAccountsChanged(accounts);
                }
            } catch (error) {
                console.error("Error checking accounts:", error);
            }
            
            // Listen for account changes
            window.ethereum.on('accountsChanged', this.handleAccountsChanged.bind(this));
            window.ethereum.on('chainChanged', () => window.location.reload());
        }
        
        // Show home page initially
        this.showPage('home');
    }
    
    setupNavigation() {
        // Handle page navigation
        document.querySelectorAll('[data-page]').forEach(link => {
            link.addEventListener('click', (e) => {
                e.preventDefault();
                const targetPage = e.target.getAttribute('data-page');
                this.showPage(targetPage);
            });
        });
    }
    
    showPage(pageId) {
        // Hide all pages
        document.querySelectorAll('.page-section').forEach(page => {
            page.classList.add('d-none');
        });
        
        // Show the requested page
        const targetPage = document.getElementById(`${pageId}-page`);
        if (targetPage) {
            targetPage.classList.remove('d-none');
            
            // Load data for the page
            if (pageId === 'oracles') {
                this.loadOracles();
            } else if (pageId === 'bets') {
                this.loadBets();
            }
        }
        
        // Update active nav link
        document.querySelectorAll('.nav-link').forEach(link => {
            link.classList.remove('active');
            if (link.getAttribute('data-page') === pageId) {
                link.classList.add('active');
            }
        });
    }
    
    setupWalletConnection() {
        const connectWalletBtn = document.getElementById('connectWalletBtn');
        if (connectWalletBtn) {
            connectWalletBtn.addEventListener('click', () => {
                this.connectWallet();
            });
        }
    }
    
    async connectWallet() {
        if (!window.ethereum) {
            this.showToast('No wallet detected. Please install MetaMask.', 'warning');
            return;
        }
        
        try {
            const accounts = await window.ethereum.request({ 
                method: 'eth_requestAccounts' 
            });
            this.handleAccountsChanged(accounts);
        } catch (error) {
            console.error("Error connecting wallet:", error);
            this.showToast('Error connecting to wallet', 'danger');
        }
    }
    
    handleAccountsChanged(accounts) {
        if (accounts.length === 0) {
            this.currentAccount = null;
            this.updateWalletStatus('Not connected');
            document.getElementById('connectWalletBtn').textContent = 'Connect Wallet';
        } else {
            this.currentAccount = accounts[0];
            this.accounts = accounts;
            
            // Initialize Web3
            this.web3 = new Web3(window.ethereum);
            
            // Update UI
            const truncatedAddress = `${accounts[0].substring(0, 6)}...${accounts[0].substring(accounts[0].length - 4)}`;
            this.updateWalletStatus(`Connected: ${truncatedAddress}`);
            document.getElementById('connectWalletBtn').textContent = 'Connected';
            
            // Load data if on specific pages
            const currentPage = document.querySelector('.page-section:not(.d-none)');
            if (currentPage) {
                const pageId = currentPage.id.replace('-page', '');
                if (pageId === 'oracles') {
                    this.loadOracles();
                } else if (pageId === 'bets') {
                    this.loadBets();
                }
            }
        }
    }
    
    updateWalletStatus(message) {
        const statusElement = document.querySelector('.wallet-status');
        if (statusElement) {
            statusElement.textContent = message;
            if (this.currentAccount) {
                statusElement.classList.add('wallet-connected');
            } else {
                statusElement.classList.remove('wallet-connected');
            }
        }
    }
    
    setupForms() {
        // Oracle creation form
        const createOracleForm = document.getElementById('createOracleForm');
        if (createOracleForm) {
            createOracleForm.addEventListener('submit', (e) => {
                e.preventDefault();
                this.createOracle();
            });
        }
        
        // Bet creation form
        const createBetForm = document.getElementById('createBetForm');
        if (createBetForm) {
            createBetForm.addEventListener('submit', (e) => {
                e.preventDefault();
                this.createBet();
            });
        }
    }
    
    async loadOracles() {
        const oraclesList = document.getElementById('oraclesList');
        if (!oraclesList) return;
        
        oraclesList.innerHTML = `
            <div class="col-12 text-center py-5">
                <div class="spinner-border" role="status">
                    <span class="visually-hidden">Loading...</span>
                </div>
            </div>
        `;
        
        try {
            const response = await fetch(`${this.apiBaseUrl}/oracles`);
            const data = await response.json();
            
            if (data.success) {
                this.deployedOracles = data.data || [];
                this.renderOracles();
            } else {
                oraclesList.innerHTML = `
                    <div class="col-12 text-center py-5">
                        <p>Failed to load oracles: ${data.error || 'Unknown error'}</p>
                    </div>
                `;
            }
        } catch (error) {
            console.error("Error loading oracles:", error);
            oraclesList.innerHTML = `
                <div class="col-12 text-center py-5">
                    <p>Error loading oracles. Please check your connection.</p>
                </div>
            `;
        }
    }
    
    renderOracles() {
        const oraclesList = document.getElementById('oraclesList');
        if (!oraclesList) return;
        
        if (this.deployedOracles.length === 0) {
            oraclesList.innerHTML = `
                <div class="col-12 text-center py-5">
                    <p>No oracles found. Create one using the form above!</p>
                </div>
            `;
            return;
        }
        
        oraclesList.innerHTML = this.deployedOracles.map(oracle => {
            let resultText = 'Undefined';
            let resultClass = 'result-undefined';
            
            if (oracle.resultSet) {
                resultText = oracle.result ? 'TRUE' : 'FALSE';
                resultClass = oracle.result ? 'result-true' : 'result-false';
            }
            
            return `
                <div class="col-md-6 col-lg-4">
                    <div class="card oracle-card" data-address="${oracle.address}">
                        <div class="card-header">
                            Oracle
                        </div>
                        <div class="card-body">
                            <p class="card-text">
                                <strong>Address:</strong><br>
                                <small>${oracle.address}</small>
                            </p>
                            <p class="card-text">
                                <strong>Trusted Wallets:</strong><br>
                                <small>${oracle.trustedWallets.join('<br>')}</small>
                            </p>
                            <p class="card-text">
                                <strong>Result:</strong> <span class="${resultClass}">${resultText}</span>
                            </p>
                            ${!oracle.resultSet && this.isTrustedWallet(oracle.address) ? `
                                <div class="d-flex gap-2 mt-3">
                                    <button class="btn btn-success btn-sm set-result-btn" 
                                        data-address="${oracle.address}" data-result="true">
                                        Set TRUE
                                    </button>
                                    <button class="btn btn-danger btn-sm set-result-btn" 
                                        data-address="${oracle.address}" data-result="false">
                                        Set FALSE
                                    </button>
                                </div>
                            ` : ''}
                        </div>
                    </div>
                </div>
            `;
        }).join('');
        
        // Add event listeners for setting results
        document.querySelectorAll('.set-result-btn').forEach(btn => {
            btn.addEventListener('click', (e) => {
                const address = e.target.getAttribute('data-address');
                const result = e.target.getAttribute('data-result') === 'true';
                this.setOracleResult(address, result);
            });
        });
    }
    
    async loadBets() {
        const betsList = document.getElementById('betsList');
        if (!betsList) return;
        
        betsList.innerHTML = `
            <div class="col-12 text-center py-5">
                <div class="spinner-border" role="status">
                    <span class="visually-hidden">Loading...</span>
                </div>
            </div>
        `;
        
        try {
            const response = await fetch(`${this.apiBaseUrl}/bets`);
            const data = await response.json();
            
            if (data.success) {
                this.deployedBets = data.data || [];
                this.renderBets();
            } else {
                betsList.innerHTML = `
                    <div class="col-12 text-center py-5">
                        <p>Failed to load bets: ${data.error || 'Unknown error'}</p>
                    </div>
                `;
            }
        } catch (error) {
            console.error("Error loading bets:", error);
            betsList.innerHTML = `
                <div class="col-12 text-center py-5">
                    <p>Error loading bets. Please check your connection.</p>
                </div>
            `;
        }
    }
    
    renderBets() {
        const betsList = document.getElementById('betsList');
        if (!betsList) return;
        
        if (this.deployedBets.length === 0) {
            betsList.innerHTML = `
                <div class="col-12 text-center py-5">
                    <p>No prediction markets found. Create one using the form above!</p>
                </div>
            `;
            return;
        }
        
        betsList.innerHTML = this.deployedBets.map(bet => {
            const releaseDate = new Date(bet.releaseDate);
            const now = new Date();
            const isReleaseReached = releaseDate <= now;
            
            return `
                <div class="col-md-6">
                    <div class="card bet-card" data-address="${bet.address}">
                        <div class="card-header d-flex justify-content-between align-items-center">
                            <div>Prediction Market</div>
                            ${bet.isClosed ? '<span class="badge bg-secondary">Closed</span>' : 
                              isReleaseReached ? '<span class="badge bg-warning text-dark">Release Date Reached</span>' : 
                              '<span class="badge bg-success">Active</span>'}
                        </div>
                        <div class="card-body">
                            <p class="bet-question">${bet.question}</p>
                            
                            <div class="bet-info">
                                <p><strong>Oracle:</strong> <small>${bet.oracleAddress}</small></p>
                                <p><strong>Release Date:</strong> ${releaseDate.toLocaleString()}</p>
                                
                                <strong>Current Odds:</strong>
                                <div class="position-relative mt-2">
                                    <div class="progress">
                                        <div class="progress-bar true-progress" role="progressbar" 
                                             style="width: ${bet.trueOdds}%" 
                                             aria-valuenow="${bet.trueOdds}" aria-valuemin="0" aria-valuemax="100">
                                        </div>
                                        <div class="progress-bar false-progress" role="progressbar" 
                                             style="width: ${bet.falseOdds}%" 
                                             aria-valuenow="${bet.falseOdds}" aria-valuemin="0" aria-valuemax="100">
                                        </div>
                                    </div>
                                    <div class="progress-label">TRUE ${bet.trueOdds}% / ${bet.falseOdds}% FALSE</div>
                                </div>
                                
                                <div class="mt-3">
                                    <p><strong>Total TRUE:</strong> ${bet.totalTrueBets} ETH</p>
                                    <p><strong>Total FALSE:</strong> ${bet.totalFalseBets} ETH</p>
                                </div>
                                
                                ${bet.resultSet ? `
                                    <div class="alert ${bet.result ? 'alert-success' : 'alert-danger'} mt-3">
                                        Result: <strong>${bet.result ? 'TRUE' : 'FALSE'}</strong>
                                    </div>
                                ` : ''}
                                
                                <div class="mt-3">
                                    ${!bet.isClosed && this.currentAccount ? `
                                        ${!isReleaseReached ? `
                                            <div class="mb-3 bet-amount-input">
                                                <label class="form-label">Bet Amount (ETH)</label>
                                                <input type="number" class="form-control bet-amount" min="0.001" step="0.001" placeholder="0.01">
                                            </div>
                                            <div class="bet-buttons">
                                                <button class="btn btn-success place-bet-btn" data-address="${bet.address}" data-prediction="true">
                                                    Bet TRUE
                                                </button>
                                                <button class="btn btn-danger place-bet-btn" data-address="${bet.address}" data-prediction="false">
                                                    Bet FALSE
                                                </button>
                                            </div>
                                        ` : isReleaseReached && !bet.isClosed && bet.resultSet ? `
                                            <button class="btn btn-primary close-bet-btn" data-address="${bet.address}">
                                                Close Bet & Distribute Funds
                                            </button>
                                        ` : ''}
                                    ` : ''}
                                </div>
                            </div>
                        </div>
                        <div class="card-footer">
                            <button class="btn btn-sm btn-outline-primary view-details-btn" data-address="${bet.address}">
                                View Details
                            </button>
                        </div>
                    </div>
                </div>
            `;
        }).join('');
        
        // Add event listeners
        document.querySelectorAll('.place-bet-btn').forEach(btn => {
            btn.addEventListener('click', (e) => {
                const address = e.target.getAttribute('data-address');
                const prediction = e.target.getAttribute('data-prediction') === 'true';
                const card = e.target.closest('.bet-card');
                const amountInput = card.querySelector('.bet-amount');
                const amount = amountInput ? amountInput.value : '0.01';
                
                this.placeBet(address, prediction, amount);
            });
        });
        
        document.querySelectorAll('.close-bet-btn').forEach(btn => {
            btn.addEventListener('click', (e) => {
                const address = e.target.getAttribute('data-address');
                this.closeBet(address);
            });
        });
        
        document.querySelectorAll('.view-details-btn').forEach(btn => {
            btn.addEventListener('click', (e) => {
                const address = e.target.getAttribute('data-address');
                this.showBetDetails(address);
            });
        });
    }
    
    isTrustedWallet(oracleAddress) {
        // Check if current wallet is trusted for this oracle
        const oracle = this.deployedOracles.find(o => o.address === oracleAddress);
        return oracle && this.currentAccount && oracle.trustedWallets.includes(this.currentAccount.toLowerCase());
    }
    
    // Create Oracle
    async createOracle() {
        if (!this.currentAccount) {
            this.showToast('Please connect your wallet first', 'warning');
            return;
        }
        
        const trustedWalletsInput = document.getElementById('trustedWallets');
        if (!trustedWalletsInput || !trustedWalletsInput.value.trim()) {
            this.showToast('Please provide at least one trusted wallet address', 'warning');
            return;
        }
        
        // Parse wallet addresses
        const trustedWallets = trustedWalletsInput.value
            .split(',')
            .map(address => address.trim())
            .filter(address => address.length > 0);
        
        if (trustedWallets.length === 0) {
            this.showToast('Please provide at least one valid wallet address', 'warning');
            return;
        }
        
        try {
            this.showToast('Deploying Oracle contract...', 'info');
            
            // Call the API to deploy oracle
            const response = await fetch(`${this.apiBaseUrl}/oracles`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    trustedWallets,
                }),
            });
            
            const result = await response.json();
            
            if (result.success) {
                this.showToast('Oracle deployed successfully!', 'success');
                trustedWalletsInput.value = '';
                this.loadOracles();
            } else {
                this.showToast(`Failed to deploy oracle: ${result.error || 'Unknown error'}`, 'danger');
            }
        } catch (error) {
            console.error("Error deploying oracle:", error);
            this.showToast('Error deploying oracle. Check console for details.', 'danger');
        }
    }
    
    // Create Bet
    async createBet() {
        if (!this.currentAccount) {
            this.showToast('Please connect your wallet first', 'warning');
            return;
        }
        
        const oracleAddress = document.getElementById('oracleAddress').value;
        const question = document.getElementById('betQuestion').value;
        const releaseDate = document.getElementById('releaseDate').value;
        
        if (!oracleAddress || !question || !releaseDate) {
            this.showToast('Please fill all required fields', 'warning');
            return;
        }
        
        // Convert release date to timestamp
        const releaseDateTimestamp = Math.floor(new Date(releaseDate).getTime() / 1000);
        
        if (releaseDateTimestamp <= Math.floor(Date.now() / 1000)) {
            this.showToast('Release date must be in the future', 'warning');
            return;
        }
        
        try {
            this.showToast('Deploying Bet contract...', 'info');
            
            // Call the API to deploy bet
            const response = await fetch(`${this.apiBaseUrl}/bets`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    oracleAddress,
                    question,
                    releaseDate: releaseDateTimestamp,
                }),
            });
            
            const result = await response.json();
            
            if (result.success) {
                this.showToast('Bet deployed successfully!', 'success');
                document.getElementById('oracleAddress').value = '';
                document.getElementById('betQuestion').value = '';
                document.getElementById('releaseDate').value = '';
                this.loadBets();
            } else {
                this.showToast(`Failed to deploy bet: ${result.error || 'Unknown error'}`, 'danger');
            }
        } catch (error) {
            console.error("Error deploying bet:", error);
            this.showToast('Error deploying bet. Check console for details.', 'danger');
        }
    }
    
    // Set Oracle Result
    async setOracleResult(oracleAddress, result) {
        if (!this.currentAccount) {
            this.showToast('Please connect your wallet first', 'warning');
            return;
        }
        
        try {
            this.showToast(`Setting oracle result to ${result ? 'TRUE' : 'FALSE'}...`, 'info');
            
            // Call the API to set result
            const response = await fetch(`${this.apiBaseUrl}/oracles/${oracleAddress}/set-result`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    result,
                }),
            });
            
            const responseData = await response.json();
            
            if (responseData.success) {
                this.showToast('Oracle result set successfully!', 'success');
                this.loadOracles();
                this.loadBets(); // Reload bets as they may be affected by the oracle result
            } else {
                this.showToast(`Failed to set oracle result: ${responseData.error || 'Unknown error'}`, 'danger');
            }
        } catch (error) {
            console.error("Error setting oracle result:", error);
            this.showToast('Error setting oracle result. Check console for details.', 'danger');
        }
    }
    
    // Place a Bet
    async placeBet(betAddress, prediction, amount) {
        if (!this.currentAccount) {
            this.showToast('Please connect your wallet first', 'warning');
            return;
        }
        
        if (!amount || parseFloat(amount) <= 0) {
            this.showToast('Please enter a valid bet amount', 'warning');
            return;
        }
        
        try {
            const amountWei = this.web3.utils.toWei(amount, 'ether');
            const betContract = new this.web3.eth.Contract(TinyBetABI, betAddress);
            
            this.showToast(`Placing ${prediction ? 'TRUE' : 'FALSE'} bet of ${amount} ETH...`, 'info');
            
            // Call the appropriate method based on prediction
            const method = prediction ? betContract.methods.betTrue() : betContract.methods.betFalse();
            
            await method.send({
                from: this.currentAccount,
                value: amountWei,
                gas: 200000
            });
            
            this.showToast('Bet placed successfully!', 'success');
            this.loadBets();
        } catch (error) {
            console.error("Error placing bet:", error);
            this.showToast('Error placing bet. Check console for details.', 'danger');
        }
    }
    
    // Close Bet
    async closeBet(betAddress) {
        if (!this.currentAccount) {
            this.showToast('Please connect your wallet first', 'warning');
            return;
        }
        
        try {
            this.showToast('Closing bet and distributing funds...', 'info');
            
            // Call the API to close bet
            const response = await fetch(`${this.apiBaseUrl}/bets/${betAddress}/close`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
            });
            
            const responseData = await response.json();
            
            if (responseData.success) {
                this.showToast('Bet closed successfully!', 'success');
                this.loadBets();
            } else {
                this.showToast(`Failed to close bet: ${responseData.error || 'Unknown error'}`, 'danger');
            }
        } catch (error) {
            console.error("Error closing bet:", error);
            this.showToast('Error closing bet. Check console for details.', 'danger');
        }
    }
    
    // Show Bet Details
    showBetDetails(betAddress) {
        const bet = this.deployedBets.find(b => b.address === betAddress);
        if (!bet) return;
        
        const releaseDate = new Date(bet.releaseDate);
        const now = new Date();
        const isReleaseReached = releaseDate <= now;
        
        // Create the modal content
        const modalTitle = document.getElementById('betDetailTitle');
        const modalBody = document.getElementById('betDetailBody');
        
        modalTitle.textContent = bet.question;
        
        modalBody.innerHTML = `
            <div class="row">
                <div class="col-md-6">
                    <h5>Bet Information</h5>
                    <p><strong>Address:</strong> <small>${bet.address}</small></p>
                    <p><strong>Oracle:</strong> <small>${bet.oracleAddress}</small></p>
                    <p><strong>Release Date:</strong> ${releaseDate.toLocaleString()}</p>
                    <p><strong>Status:</strong> 
                        ${bet.isClosed ? '<span class="badge bg-secondary">Closed</span>' : 
                          isReleaseReached ? '<span class="badge bg-warning text-dark">Release Date Reached</span>' : 
                          '<span class="badge bg-success">Active</span>'}
                    </p>
                    
                    ${bet.resultSet ? `
                        <div class="alert ${bet.result ? 'alert-success' : 'alert-danger'} mt-3">
                            Result: <strong>${bet.result ? 'TRUE' : 'FALSE'}</strong>
                        </div>
                    ` : ''}
                </div>
                
                <div class="col-md-6">
                    <h5>Betting Statistics</h5>
                    <div class="odds-container">
                        <div class="odds-chart-container">
                            <canvas id="oddsChart"></canvas>
                        </div>
                    </div>
                    <div class="mt-3">
                        <p><strong>Total TRUE:</strong> ${bet.totalTrueBets} ETH</p>
                        <p><strong>Total FALSE:</strong> ${bet.totalFalseBets} ETH</p>
                    </div>
                </div>
            </div>
            
            ${!bet.isClosed && this.currentAccount ? `
                <hr>
                <div class="row mt-3">
                    <div class="col-12">
                        ${!isReleaseReached ? `
                            <h5>Place a Bet</h5>
                            <div class="row align-items-center">
                                <div class="col-md-4">
                                    <div class="mb-3">
                                        <label class="form-label">Bet Amount (ETH)</label>
                                        <input type="number" class="form-control bet-amount-detail" min="0.001" step="0.001" placeholder="0.01">
                                    </div>
                                </div>
                                <div class="col-md-8">
                                    <div class="bet-buttons">
                                        <button class="btn btn-success place-bet-detail-btn" data-address="${bet.address}" data-prediction="true">
                                            Bet TRUE
                                        </button>
                                        <button class="btn btn-danger place-bet-detail-btn" data-address="${bet.address}" data-prediction="false">
                                            Bet FALSE
                                        </button>
                                    </div>
                                </div>
                            </div>
                        ` : isReleaseReached && !bet.isClosed && bet.resultSet ? `
                            <div class="text-center">
                                <button class="btn btn-primary close-bet-detail-btn" data-address="${bet.address}">
                                    Close Bet & Distribute Funds
                                </button>
                            </div>
                        ` : ''}
                    </div>
                </div>
            ` : ''}
        `;
        
        // Show the modal
        const betDetailModal = new bootstrap.Modal(document.getElementById('betDetailModal'));
        betDetailModal.show();
        
        // Initialize chart if odds data exists
        if (bet.trueOdds !== undefined && bet.falseOdds !== undefined) {
            setTimeout(() => {
                const ctx = document.getElementById('oddsChart').getContext('2d');
                new Chart(ctx, {
                    type: 'pie',
                    data: {
                        labels: ['TRUE', 'FALSE'],
                        datasets: [{
                            data: [bet.trueOdds, bet.falseOdds],
                            backgroundColor: ['#198754', '#dc3545'],
                            borderWidth: 1
                        }]
                    },
                    options: {
                        responsive: true,
                        plugins: {
                            legend: {
                                position: 'bottom'
                            }
                        }
                    }
                });
            }, 100);
        }
        
        // Add event listeners for detail view
        setTimeout(() => {
            document.querySelectorAll('.place-bet-detail-btn').forEach(btn => {
                btn.addEventListener('click', (e) => {
                    const address = e.target.getAttribute('data-address');
                    const prediction = e.target.getAttribute('data-prediction') === 'true';
                    const amountInput = document.querySelector('.bet-amount-detail');
                    const amount = amountInput ? amountInput.value : '0.01';
                    
                    this.placeBet(address, prediction, amount);
                    betDetailModal.hide();
                });
            });
            
            document.querySelectorAll('.close-bet-detail-btn').forEach(btn => {
                btn.addEventListener('click', (e) => {
                    const address = e.target.getAttribute('data-address');
                    this.closeBet(address);
                    betDetailModal.hide();
                });
            });
        }, 100);
    }
    
    // Show toast notification
    showToast(message, type = 'info') {
        // Create toast container if it doesn't exist
        let toastContainer = document.querySelector('.toast-container');
        if (!toastContainer) {
            toastContainer = document.createElement('div');
            toastContainer.className = 'toast-container';
            document.body.appendChild(toastContainer);
        }
        
        // Create the toast
        const toastId = 'toast-' + Date.now();
        const toast = document.createElement('div');
        toast.className = `toast align-items-center text-white bg-${type} border-0`;
        toast.setAttribute('role', 'alert');
        toast.setAttribute('aria-live', 'assertive');
        toast.setAttribute('aria-atomic', 'true');
        toast.setAttribute('id', toastId);
        
        toast.innerHTML = `
            <div class="d-flex">
                <div class="toast-body">
                    ${message}
                </div>
                <button type="button" class="btn-close btn-close-white me-2 m-auto" data-bs-dismiss="toast" aria-label="Close"></button>
            </div>
        `;
        
        toastContainer.appendChild(toast);
        
        // Show the toast
        const bsToast = new bootstrap.Toast(toast, {
            autohide: true,
            delay: 5000
        });
        bsToast.show();
        
        // Remove toast when hidden
        toast.addEventListener('hidden.bs.toast', () => {
            toast.remove();
        });
    }
}

// Initialize the app when DOM is loaded
document.addEventListener('DOMContentLoaded', () => {
    const app = new PredictionMarketApp();
});