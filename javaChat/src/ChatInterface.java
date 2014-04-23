import java.awt.BorderLayout;
import java.awt.Color;
import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;
import java.io.IOException;

import javax.swing.*;


public class ChatInterface implements ActionListener {

	JTextField field;
	
	Testclientsocket client; 
	JTextArea chatPanel;
	String message ="";
	boolean reconnect = true;
	
	public static void main(String[] args) {
	
		ChatInterface ci = new ChatInterface();
		ci.createChatGUI();
		
	}
	
	
	/*
	 * Creates a simple chat GUI
	 */
	public void createChatGUI(){
				
		JFrame chat = new JFrame();
		chat.setLayout(new BorderLayout());
		chat.getContentPane().setBackground(Color.LIGHT_GRAY);
		chat.setSize(800,600);
        chat.setLocationRelativeTo(null);
        chat.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
        chat.setVisible(true);
        

        field = new JTextField(20);
        
        
        chatPanel = new JTextArea();
        chatPanel.setVisible(true);
        chatPanel.setEditable(false);

        chat.add(chatPanel, "North");
        chat.add(field, "South");
        
        testConnection();
		
        
		field.addActionListener(this);
		updateChat();

	}
	/*
	 * Updates the textChat by receiving message from the server, Handles if a connection failed
	 */
	public void updateChat(){
    	while(true){
    		
    		try {
				message = message + "\n" + client.recieveMessage();
			} catch (IOException e) {
				testConnection();
			}
       		chatPanel.setText(message);
    		
    	}
		
	}
	
	/*
	 * Test the connection, if failed lets the user try to connect again 
	 *
	 */
	public void testConnection(){
        while(true){
		try {
			client = new Testclientsocket();
			return;
		} catch (IOException e) {
			int reply = JOptionPane.showConfirmDialog(null,"Server is down! Would you like to reconnect?", "Warning", JOptionPane.YES_NO_OPTION); 
			if (reply == JOptionPane.NO_OPTION)
				System.exit(1);
		}
			
		}
		
		
	}


	
	@Override
	public void actionPerformed(ActionEvent arg0) {
		String text = field.getText();
		client.sendMessage(text);
		// TODO Auto-generated method stub
		
	}



	
	
}
