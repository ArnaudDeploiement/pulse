Outil CLI en Go permettant de créer et gérer des groupes privés pour le partage direct de fichiers via le réseau libp2p, sans serveur central et sans stockage intermédiaire. Pulse utilise un relay passif uniquement pour contourner les NAT/pare-feux, et ne conserve aucune donnée sur disque en dehors des machines des membres.

✨ Fonctionnalités
🔐 Groupes privés avec identité basée sur clé privée Ed25519

📄 Fichiers GroupKey.json (privé) et Invite.json (public) générés automatiquement

🌍 Relay passif libp2p pour permettre la connexion derrière NAT

⏱ Transferts temps réel uniquement (aucun stockage sur le relay)

⚡ Connexion directe si possible, fallback via relay sinon

🛠 Interface CLI simple (basée sur Cobra)


📜 Licence
MIT — utilisation libre avec attribution.


