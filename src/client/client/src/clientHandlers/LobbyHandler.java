package clientHandlers;

import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;
import java.io.IOException;

import javax.swing.JOptionPane;

import org.json.JSONException;
import org.json.JSONObject;

import packageManaging.CreateGameMessage;
import packageManaging.CreateGameMessageEncoder;
import packageManaging.JoinMessage;
import packageManaging.JoinMessageEncoder;
import packageManaging.LobbyMessageEncoder;
import packageManaging.LobbyServerMessage;
import packageManaging.Message;
import packageManaging.ChatMessageEncoder;
import packageManaging.RefeshMessageEncoder;
import packageManaging.RefreshMessage;
import packageManaging.StageFlipper;
import graphicalReference.ChatGUI;
import graphicalReference.LobbyGUI;

import clientCore.Monitor;
import clientNetworking.NetManager;

public class LobbyHandler extends Monitor implements HandlerInterface, ActionListener {
	LobbyGUI lg;
	ChatGUI cg;
	LobbyMessageEncoder lme = new LobbyMessageEncoder();
	ChatMessageEncoder cme = new ChatMessageEncoder();
	JoinMessageEncoder jme = new JoinMessageEncoder();
	CreateGameMessageEncoder gme = new CreateGameMessageEncoder();
	RefeshMessageEncoder  rme = new RefeshMessageEncoder();
	NetManager network;
	public boolean loop;
	int state;
	final String userName;
	
	public LobbyHandler(NetManager net, String username){
		this.network = net;
		this.userName = username;
		this.loop = true;
	}
	
	public void init() {

		String firstcall;
		LobbyServerMessage LM = null;
		try {
			firstcall = network.receiveMessage();
			LM = lme.decode(firstcall);
			
		} catch (IOException | JSONException e) {
			e.printStackTrace();
		}

		if (LM != null) {
			lg = new LobbyGUI(this.userName);
			lg.render(LM.UserList, LM.GameHost);
			lg.chatgui.field.addActionListener(this);
			lg.joinButton.addActionListener(this);
			lg.createButton.addActionListener(this);
			lg.refreshButton.addActionListener(this);
			
		} else {
			System.out.println("Connection failed");
		}
		runLobby();
	}
	
	public void runLobby() {
		if (super.lastMsg != null) {
			lg.lobby.setVisible(true);
			StageFlipper startup = super.lastMsg;
		}
		while (loop) {
			try {
				String mess = receiveMessage();
			
				int id = retrieveId(mess);
				decodeAndRender(id, mess);
				
			} catch (IOException | JSONException e) {
				e.printStackTrace();
			}
		}
		lg.lobby.setVisible(false);
	}

	public void actionPerformed(ActionEvent e) {
		String command = e.getActionCommand();
		String JSONtext = "";
		switch (command) {
		case "chatmessage":  //JTextField
			try {
				JSONtext = cme.encode(new Message(lg.chatgui.field.getText(), ""));
				lg.chatgui.field.setText("");
			} catch (JSONException e1) {
				e1.printStackTrace();
			}
			break;

		case "joinbutton":
			try {
				String players = (String) lg.jt.getModel().getValueAt(lg.jt.getSelectedRow(),1);
				
				if (!(players.charAt(0) == players.charAt(2))) {
				JSONtext = jme.encode(new JoinMessage((String)lg.jt.getModel().getValueAt(lg.jt.getSelectedRow(),3)));
				} else {
					JOptionPane.showMessageDialog(null, "Room is full!", "Warning", JOptionPane.ERROR_MESSAGE);
				}
			} catch (JSONException e1) {
				// TODO Auto-generated catch block
				e1.printStackTrace();
			}
			break;
			
		case "createbutton":
			try {
				JSONtext = gme.encode(new CreateGameMessage(lg.createList.getSelectedValue(), 2 , this.userName));
				sendMessage(JSONtext);
				JSONtext = rme.encode(new RefreshMessage());
				
			} catch (JSONException e1) {
				// TODO Auto-generated catch block
				e1.printStackTrace();
			}
			break;
			
		case "refreshbutton":		
			try {
				JSONtext = rme.encode(new RefreshMessage());
			} catch (JSONException e1) {
				// TODO Auto-generated catch block
				e1.printStackTrace();
			}
			break;
			
		default:
			JSONtext = "";

		}
		sendMessage(JSONtext);

	}

	public void decodeAndRender(int id, String message) throws JSONException {
		switch (id) {
		case 100: // Message.java
			Message chatMessage = cme.decode(message);
			lg.chatgui.chatUpdate(chatMessage.message, chatMessage.user);
			break;

		case 101: // Create session

		case 102: // Join session
			StageFlipper flip = new StageFlipper(jme.decode(message));
			super.stop(flip);
			break;

		case 103: // LobbyClientMessage.java/LobbyServerMessage.java
			LobbyServerMessage ls = lme.decode(message);
			lg.updateJTable(ls.UserList);
			break;
		default:
			System.out.println("something went wrong! Error code: " + id + "\nMessage: " + message);
		}

	}

	public void sendMessage(String message) {
		network.send(message);
	}

	public String receiveMessage() throws IOException {
		return network.receiveMessage();

	}

	public int retrieveId(String mess) {
		try {
			JSONObject obj = new JSONObject(mess);
			return obj.getInt("PacketID");
		} catch (JSONException e) {
			
		}
		return 0;

	}
}
